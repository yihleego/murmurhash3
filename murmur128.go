package murmur3

import "math/bits"

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

func (h *MurmurHash128) HashInt(i int) *Int32HashCode {
	if bits.UintSize == 32 {
		return h.HashInt32(int32(i))
	}
	return h.HashInt64(int64(i))
}

func (h *MurmurHash128) HashInt32(i int32) *Int32HashCode {
	k1 := mixK1(uint32(i))
	h1 := mixH1(uint32(h.seed), k1)
	return &Int32HashCode{int32(mixFinal(h1, 4))}
}

func (h *MurmurHash128) HashInt64(i int64) *Int32HashCode {
	low := uint32(i)
	high := uint32(i >> 32)
	k1 := mixK1(low)
	h1 := mixH1(uint32(h.seed), k1)
	k1 = mixK1(high)
	h1 = mixH1(h1, k1)
	return &Int32HashCode{int32(mixFinal(h1, 8))}
}

func (h *MurmurHash128) HashBytes(bytes []byte, offset int, length int) *Int32HashCode {

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

func (h *MurmurHash128) HashFullBytes(bytes []byte) *Int32HashCode {
	return h.HashBytes(bytes, 0, len(bytes))
}
