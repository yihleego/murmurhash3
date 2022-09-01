package murmur3

import (
	"fmt"
	"math"
	"testing"
)

func TestHashInt(t *testing.T) {
	murmur := New32()
	ints := []int{0, 1, -1, math.MaxInt, math.MinInt}
	for _, i := range ints {
		hash := murmur.HashInt(i)
		fmt.Println(hash)
		fmt.Println(hash.AsInt())
		fmt.Println(hash.AsInt32())
		fmt.Println(hash.AsInt64())
		fmt.Println(hash.AsBytes())
	}
}

func TestHashInt32(t *testing.T) {
	murmur := New32()
	ints := []int32{0, 1, -1, math.MaxInt32, math.MinInt32}
	for _, i := range ints {
		hash := murmur.HashInt32(i)
		fmt.Println(hash)
		fmt.Println(hash.AsInt())
		fmt.Println(hash.AsInt32())
		fmt.Println(hash.AsInt64())
		fmt.Println(hash.AsBytes())
	}
}

func TestHashInt64(t *testing.T) {
	murmur := New32()
	ints := []int64{0, 1, -1, math.MaxInt64, math.MinInt64}
	for _, i := range ints {
		hash := murmur.HashInt64(i)
		fmt.Println(hash)
		fmt.Println(hash.AsInt())
		fmt.Println(hash.AsInt32())
		fmt.Println(hash.AsInt64())
		fmt.Println(hash.AsBytes())
	}
}

func TestHashBytes(t *testing.T) {
	murmur := New32()
	bytes := [][]byte{{11, 22, 33, 44}}
	for _, b := range bytes {
		hash := murmur.HashFullBytes(b)
		fmt.Println(hash)
		fmt.Println(hash.AsInt())
		fmt.Println(hash.AsInt32())
		fmt.Println(hash.AsInt64())
		fmt.Println(hash.AsBytes())
	}
}

func TestHashString(t *testing.T) {
	murmur := New32()
	strings := []string{"hello, world"}
	for _, s := range strings {
		hash := murmur.HashString(s)
		fmt.Println(hash)
		fmt.Println(hash.AsInt())
		fmt.Println(hash.AsInt32())
		fmt.Println(hash.AsInt64())
		fmt.Println(hash.AsBytes())
	}
}
