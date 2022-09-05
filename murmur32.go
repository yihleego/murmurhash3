package murmur3

import (
	"math/bits"
	"reflect"
	"unsafe"
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
	k1 := h.mixK1(uint32(i))
	h1 := h.mixH1(uint32(h.seed), k1)

	h1 = h.fmix(h1, 4)
	return h.makeHash(h1)
}

func (h *MurmurHash32) HashInt64(i int64) *Int32HashCode {
	low := uint32(i)
	high := uint32(i >> 32)

	k1 := h.mixK1(low)
	h1 := h.mixH1(uint32(h.seed), k1)

	k1 = h.mixK1(high)
	h1 = h.mixH1(h1, k1)

	h1 = h.fmix(h1, 8)
	return h.makeHash(h1)
}

func (h *MurmurHash32) HashString(s string) *Int32HashCode {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{Data: sh.Data, Len: sh.Len, Cap: sh.Len}
	bytes := *(*[]byte)(unsafe.Pointer(&bh))
	return h.HashBytes(bytes)
}

func (h *MurmurHash32) HashBytes(bytes []byte) *Int32HashCode {
	return h.make(bytes)
}

func (h *MurmurHash32) HashBytesWithOffset(bytes []byte, offset, length int) *Int32HashCode {
	return h.make(bytes[offset : offset+length])
}

func (h *MurmurHash32) make(bytes []byte) *Int32HashCode {
	length := uint32(len(bytes))
	h1, buffer := h.bmix(uint32(h.seed), bytes)

	var k1 uint32
	switch len(buffer) & 3 {
	case 3:
		k1 ^= uint32(buffer[2]) << 16
		fallthrough
	case 2:
		k1 ^= uint32(buffer[1]) << 8
		fallthrough
	case 1:
		k1 ^= uint32(buffer[0])
	}

	h1 ^= h.mixK1(k1)

	h1 = h.fmix(h1, length)

	return h.makeHash(h1)
}

func (h *MurmurHash32) mixK1(k1 uint32) uint32 {
	k1 *= 0xcc9e2d51
	k1 = bits.RotateLeft32(k1, 15)
	k1 *= 0x1b873593
	return k1
}

func (h *MurmurHash32) mixH1(h1, k1 uint32) uint32 {
	h1 ^= k1
	h1 = bits.RotateLeft32(h1, 13)
	h1 = h1*5 + 0xe6546b64
	return h1
}

func (h *MurmurHash32) bmix(h1 uint32, bytes []byte) (uint32, []byte) {
	blocks := len(bytes) / 4
	for i := 0; i < blocks; i++ {
		k1 := *(*uint32)(unsafe.Pointer(&bytes[i*4]))
		k1 = h.mixK1(k1)
		h1 = h.mixH1(h1, k1)
	}
	return h1, bytes[blocks*4:]
}

// fmix Finalization mix - force all bits of a hash block to avalanche
func (h *MurmurHash32) fmix(h1, length uint32) uint32 {
	h1 ^= length
	h1 ^= h1 >> 16
	h1 *= 0x85ebca6b
	h1 ^= h1 >> 13
	h1 *= 0xc2b2ae35
	h1 ^= h1 >> 16
	return h1
}

func (h *MurmurHash32) makeHash(h1 uint32) *Int32HashCode {
	return &Int32HashCode{int32(h1)}
}
