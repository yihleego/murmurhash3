package murmur3

import (
	"encoding/hex"
	"math/bits"
)

type HashCode interface {
	Bits() int
	AsInt() int
	AsInt32() int32
	AsInt64() int64
	AsBytes() []byte
	AsString() string
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
	l := minInt(len(h.bytes), 4)
	v := int32(h.bytes[0]) & 0xFF
	for i := 1; i < l; i++ {
		v |= (int32(h.bytes[i]) & 0xFF) << (i * 8)
	}
	return v
}

func (h *BytesHashCode) AsInt64() int64 {
	l := minInt(len(h.bytes), 8)
	v := int64(h.bytes[0]) & 0xFF
	for i := 1; i < l; i++ {
		v |= (int64(h.bytes[i]) & 0xFF) << (i * 8)
	}
	return v
}

func (h *BytesHashCode) AsBytes() []byte {
	replica := make([]byte, len(h.bytes))
	copy(h.bytes, replica)
	return replica
}

func (h *BytesHashCode) AsString() string { return hex.EncodeToString(h.AsBytes()) }

func (h *BytesHashCode) String() string { return h.AsString() }

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

func (h *Int32HashCode) AsString() string { return hex.EncodeToString(h.AsBytes()) }

func (h *Int32HashCode) String() string { return h.AsString() }

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

func (h *Int64HashCode) AsString() string { return hex.EncodeToString(h.AsBytes()) }

func (h *Int64HashCode) String() string { return h.AsString() }

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}
