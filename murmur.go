package murmur3

import (
	"encoding/hex"
	"math/bits"
)

type MurmurHash struct {
	seed int
}

type HashCode interface {
	Bits() int
	AsInt() int
	AsInt32() int32
	AsInt64() int64
	AsBytes() []byte
	AsHex() string
}

type BytesHashCode struct {
	bytes []byte
}

func (h *BytesHashCode) Bits() int { return len(h.bytes) * 8 }

func (h *BytesHashCode) AsInt() int {
	if bits.UintSize == 32 {
		return int(h.AsInt32())
	} else {
		return int(h.AsInt64())
	}
}

func (h *BytesHashCode) AsInt32() int32 {
	l := len(h.bytes)
	if l > 4 {
		l = 4
	}
	v := int32(h.bytes[0])
	for i := 1; i < l; i++ {
		v |= (int32(h.bytes[i])) << (i * 8)
	}
	return v
}

func (h *BytesHashCode) AsInt64() int64 {
	l := len(h.bytes)
	if l > 8 {
		l = 8
	}
	v := int64(h.bytes[0])
	for i := 1; i < l; i++ {
		v |= (int64(h.bytes[i])) << (i * 8)
	}
	return v
}

func (h *BytesHashCode) AsBytes() []byte {
	replica := make([]byte, len(h.bytes))
	copy(replica, h.bytes)
	return replica
}

func (h *BytesHashCode) AsHex() string { return hex.EncodeToString(h.AsBytes()) }

func (h *BytesHashCode) String() string { return h.AsHex() }

type Int32HashCode struct{ hash int32 }

func (h *Int32HashCode) Bits() int { return 32 }

func (h *Int32HashCode) AsInt() int { return int(h.hash) }

func (h *Int32HashCode) AsInt32() int32 { return h.hash }

func (h *Int32HashCode) AsInt64() int64 { return int64(h.hash) }

func (h *Int32HashCode) AsBytes() []byte {
	return []byte{
		byte(h.hash),
		byte(h.hash >> 8),
		byte(h.hash >> 16),
		byte(h.hash >> 24)}
}

func (h *Int32HashCode) AsHex() string { return hex.EncodeToString(h.AsBytes()) }

func (h *Int32HashCode) String() string { return h.AsHex() }

type Int64HashCode struct{ hash int64 }

func (h *Int64HashCode) Bits() int { return 64 }

func (h *Int64HashCode) AsInt() int { return int(h.hash) }

func (h *Int64HashCode) AsInt32() int32 { return int32(h.hash) }

func (h *Int64HashCode) AsInt64() int64 { return h.hash }

func (h *Int64HashCode) AsBytes() []byte {
	return []byte{
		byte(h.hash),
		byte(h.hash >> 8),
		byte(h.hash >> 16),
		byte(h.hash >> 24),
		byte(h.hash >> 32),
		byte(h.hash >> 40),
		byte(h.hash >> 48),
		byte(h.hash >> 56)}
}

func (h *Int64HashCode) AsHex() string { return hex.EncodeToString(h.AsBytes()) }

func (h *Int64HashCode) String() string { return h.AsHex() }
