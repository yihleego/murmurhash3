package murmur3

import (
	"math/bits"
	"reflect"
	"unsafe"
)

const (
	C1 uint32 = 0xcc9e2d51
	C2 uint32 = 0x1b873593
)

func New32() *MurmurHash32 {
	return New32WithSeed(0)
}

func New32WithSeed(seed int) *MurmurHash32 {
	h := &MurmurHash32{}
	h.seed = seed
	return h
}

type MurmurHash32 struct {
	MurmurHash
}

func (h *MurmurHash32) HashInt(i int) *Int32HashCode {
	if bits.UintSize == 32 {
		return h.HashInt32(int32(i))
	}
	return h.HashInt64(int64(i))
}

func (h *MurmurHash32) HashInt32(i int32) *Int32HashCode {
	k1 := mixK1(uint32(i))
	h1 := mixH1(uint32(h.seed), k1)
	return &Int32HashCode{int32(mixFinal(h1, 4))}
}

func (h *MurmurHash32) HashInt64(i int64) *Int32HashCode {
	low := uint32(i)
	high := uint32(i >> 32)
	k1 := mixK1(low)
	h1 := mixH1(uint32(h.seed), k1)
	k1 = mixK1(high)
	h1 = mixH1(h1, k1)
	return &Int32HashCode{int32(mixFinal(h1, 8))}
}

func (h *MurmurHash32) HashBytes(bytes []byte, offset int, length int) *Int32HashCode {

	h1 := uint32(h.seed)
	var i int
	for i = 0; i+4 <= length; i += 4 {
		k1 := mixK1(getIntLittleEndian(bytes, offset+i))
		h1 = mixH1(h1, k1)
	}
	k1 := uint32(0)
	for shift := 0; i < length; shift += 8 {
		k1 ^= uint32(bytes[offset+i]) & 0xFF << shift
		i++
	}
	h1 ^= mixK1(k1)
	return &Int32HashCode{int32(mixFinal(h1, uint32(length)))}
}

func (h *MurmurHash32) HashFullBytes(bytes []byte) *Int32HashCode {
	return h.HashBytes(bytes, 0, len(bytes))
}

func (h *MurmurHash32) HashString(s string) *Int32HashCode {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}
	bytes := *(*[]byte)(unsafe.Pointer(&bh))
	return h.HashBytes(bytes, 0, len(bytes))
}

func mixK1(k1 uint32) uint32 {
	k1 *= C1
	k1 = bits.RotateLeft32(k1, 15)
	k1 *= C2
	return k1
}

func mixH1(h1, k1 uint32) uint32 {
	h1 ^= k1
	h1 = bits.RotateLeft32(h1, 13)
	h1 = h1*5 + 0xe6546b64
	return h1
}

func mixFinal(h1, length uint32) uint32 {
	h1 ^= length
	h1 ^= h1 >> 16
	h1 *= 0x85ebca6b
	h1 ^= h1 >> 13
	h1 *= 0xc2b2ae35
	h1 ^= h1 >> 16
	return h1
}

func getIntLittleEndian(bytes []byte, offset int) uint32 {
	return uint32(bytes[offset+3])<<24 |
		(uint32(bytes[offset+2])&0xFF)<<16 |
		(uint32(bytes[offset+1])&0xFF)<<8 |
		(uint32(bytes[offset]) & 0xFF)
}
