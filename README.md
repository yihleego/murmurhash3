# MurmurHash3

MurmurHash is a non-cryptographic hash function suitable for general hash-based lookup.

The current version is MurmurHash3, which yields a 32-bit or 128-bit hash value. When using 128-bits, the x86 and x64 versions do not produce the same values, as the algorithms are optimized for their respective platforms.

## Usage

### Basic

```go
murmur := murmur3.New32()
h := murmur.HashInt(0)
h.AsInt()   // 1669671676
h.AsInt32() // 1669671676
h.AsInt64() // 1669671676
h.AsBytes() // {252, 42, 133, 99}
h.AsHex()   // fc2a8563
```

### More types

```go
murmur := murmur3.New32()
murmur.HashInt32(0)         // def96223
murmur.HashInt64(0)         // fc2a8563
murmur.HashBytes([]byte{0}) // b7284e51
murmur.HashString("0")      // 7fc071d2
```

### Specify seeds

```go
seed := 0xFF
murmur := murmur3.New32WithSeed(seed)
```

### Use 128-bit

```go
murmur := murmur3.New128()
murmur.HashInt32(0)         // bc764cd8ddf7a0cff126f51c16239658
murmur.HashInt64(0)         // cbc357ccb763df2852fee8c4fc7d55f2
murmur.HashBytes([]byte{0}) // b55cff6ee5ab10468335f878aa2d6251
murmur.HashString("0")      // 80a346d5bedec92a095e873ce5e98d3a
```

## License

This project is under the MIT license. See the [LICENSE](LICENSE) file for details.
