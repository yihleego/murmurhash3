package murmur3

import (
	"math/bits"
	"reflect"
	"unsafe"
)

func New128() *MurmurHash128 {
	return New128WithSeed(0)
}

func New128WithSeed(seed int) *MurmurHash128 {
	h := &MurmurHash128{}
	h.seed = seed
	return h
}

type MurmurHash128 struct {
	MurmurHash
}

func (h *MurmurHash128) HashInt(i int) *BytesHashCode {
	if bits.UintSize == 32 {
		return h.HashInt32(int32(i))
	}
	return h.HashInt64(int64(i))
}

func (h *MurmurHash128) HashInt32(i int32) *BytesHashCode {
	return h.make([]byte{byte(i), byte(i >> 8), byte(i >> 16), byte(i >> 24)})
}

func (h *MurmurHash128) HashInt64(i int64) *BytesHashCode {
	return h.make([]byte{
		byte(i), byte(i >> 8), byte(i >> 16), byte(i >> 24),
		byte(i >> 32), byte(i >> 40), byte(i >> 48), byte(i >> 56)})
}

func (h *MurmurHash128) HashString(s string) *BytesHashCode {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{Data: sh.Data, Len: sh.Len, Cap: sh.Len}
	bytes := *(*[]byte)(unsafe.Pointer(&bh))
	return h.make(bytes)
}

func (h *MurmurHash128) HashBytes(bytes []byte) *BytesHashCode {
	return h.make(bytes)
}

func (h *MurmurHash128) HashBytesWithOffset(bytes []byte, offset, length int) *BytesHashCode {
	return h.make(bytes[offset : offset+length])
}

func (h *MurmurHash128) make(bytes []byte) *BytesHashCode {
	length := uint64(len(bytes))
	h1, h2, buffer := h.bmix(uint64(h.seed), uint64(h.seed), bytes)

	var k1, k2 uint64
	switch len(buffer) & 15 {
	case 15:
		k2 ^= uint64(buffer[14]) << 48
		fallthrough
	case 14:
		k2 ^= uint64(buffer[13]) << 40
		fallthrough
	case 13:
		k2 ^= uint64(buffer[12]) << 32
		fallthrough
	case 12:
		k2 ^= uint64(buffer[11]) << 24
		fallthrough
	case 11:
		k2 ^= uint64(buffer[10]) << 16
		fallthrough
	case 10:
		k2 ^= uint64(buffer[9]) << 8
		fallthrough
	case 9:
		k2 ^= uint64(buffer[8])
		fallthrough
	case 8:
		k1 ^= uint64(buffer[7]) << 56
		fallthrough
	case 7:
		k1 ^= uint64(buffer[6]) << 48
		fallthrough
	case 6:
		k1 ^= uint64(buffer[5]) << 40
		fallthrough
	case 5:
		k1 ^= uint64(buffer[4]) << 32
		fallthrough
	case 4:
		k1 ^= uint64(buffer[3]) << 24
		fallthrough
	case 3:
		k1 ^= uint64(buffer[2]) << 16
		fallthrough
	case 2:
		k1 ^= uint64(buffer[1]) << 8
		fallthrough
	case 1:
		k1 ^= uint64(buffer[0])
	}

	h1 ^= h.mixK1(k1)
	h2 ^= h.mixK2(k2)

	h1 ^= length
	h2 ^= length

	h1 += h2
	h2 += h1

	h1 = h.fmix(h1)
	h2 = h.fmix(h2)

	h1 += h2
	h2 += h1

	return h.makeHash(h1, h2)
}

func (h *MurmurHash128) mixK1(k1 uint64) uint64 {
	k1 *= 0x87c37b91114253d5
	k1 = bits.RotateLeft64(k1, 31)
	k1 *= 0x4cf5ad432745937f
	return k1
}

func (h *MurmurHash128) mixK2(k2 uint64) uint64 {
	k2 *= 0x4cf5ad432745937f
	k2 = bits.RotateLeft64(k2, 33)
	k2 *= 0x87c37b91114253d5
	return k2
}

func (h *MurmurHash128) bmix(h1, h2 uint64, bytes []byte) (uint64, uint64, []byte) {
	blocks := len(bytes) / 16
	for i := 0; i < blocks; i++ {
		t := (*[2]uint64)(unsafe.Pointer(&bytes[i*16]))
		k1, k2 := t[0], t[1]

		h1 ^= h.mixK1(k1)

		h1 = bits.RotateLeft64(h1, 27)
		h1 += h2
		h1 = h1*5 + 0x52dce729

		h2 ^= h.mixK2(k2)

		h2 = bits.RotateLeft64(h2, 31)
		h2 += h1
		h2 = h2*5 + 0x38495ab5
	}
	return h1, h2, bytes[blocks*16:]
}

func (h *MurmurHash128) fmix(k uint64) uint64 {
	k ^= k >> 33
	k *= 0xff51afd7ed558ccd
	k ^= k >> 33
	k *= 0xc4ceb9fe1a85ec53
	k ^= k >> 33
	return k
}

func (h *MurmurHash128) makeHash(h1, h2 uint64) *BytesHashCode {
	return &BytesHashCode{[]byte{
		byte(h1 >> 0), byte(h1 >> 8), byte(h1 >> 16), byte(h1 >> 24),
		byte(h1 >> 32), byte(h1 >> 40), byte(h1 >> 48), byte(h1 >> 56),
		byte(h2 >> 0), byte(h2 >> 8), byte(h2 >> 16), byte(h2 >> 24),
		byte(h2 >> 32), byte(h2 >> 40), byte(h2 >> 48), byte(h2 >> 56)}}
}
