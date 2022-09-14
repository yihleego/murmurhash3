# MurmurHash3

[![GoDoc](https://godoc.org/github.com/yihleego/murmurhash3?status.svg)](https://godoc.org/github.com/yihleego/murmurhash3)
[![Go Report Card](https://goreportcard.com/badge/github.com/yihleego/murmurhash3)](https://goreportcard.com/report/github.com/yihleego/murmurhash3)

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
## Benchmark

```text
BenchmarkMurmur32HashInt
BenchmarkMurmur32HashInt-8       	98502462	        11.86 ns/op	       4 B/op	       1 allocs/op
BenchmarkMurmur32HashInt32
BenchmarkMurmur32HashInt32-8     	92524768	        12.21 ns/op	       4 B/op	       1 allocs/op
BenchmarkMurmur32HashInt64
BenchmarkMurmur32HashInt64-8     	100000000	        11.94 ns/op	       4 B/op	       1 allocs/op
BenchmarkMurmur32HashBytes
BenchmarkMurmur32HashBytes/1
BenchmarkMurmur32HashBytes/1-8   	85952496	        13.71 ns/op	  72.94 MB/s	       4 B/op	       1 allocs/op
BenchmarkMurmur32HashBytes/4
BenchmarkMurmur32HashBytes/4-8   	85774940	        14.24 ns/op	 280.97 MB/s	       4 B/op	       1 allocs/op
BenchmarkMurmur32HashBytes/16
BenchmarkMurmur32HashBytes/16-8  	70771407	        17.13 ns/op	 934.26 MB/s	       4 B/op	       1 allocs/op
BenchmarkMurmur32HashBytes/64
BenchmarkMurmur32HashBytes/64-8  	40107621	        29.69 ns/op	2155.67 MB/s	       4 B/op	       1 allocs/op
BenchmarkMurmur32HashBytes/256
BenchmarkMurmur32HashBytes/256-8 	14972842	        79.46 ns/op	3221.55 MB/s	       4 B/op	       1 allocs/op
BenchmarkMurmur32HashBytes/1024
BenchmarkMurmur32HashBytes/1024-8         	 4247964	       282.1 ns/op	3629.33 MB/s	       4 B/op	       1 allocs/op
BenchmarkMurmur32HashString
BenchmarkMurmur32HashString/1
BenchmarkMurmur32HashString/1-8           	80060311	        14.79 ns/op	  67.63 MB/s	       4 B/op	       1 allocs/op
BenchmarkMurmur32HashString/4
BenchmarkMurmur32HashString/4-8           	77204179	        16.44 ns/op	 243.37 MB/s	       4 B/op	       1 allocs/op
BenchmarkMurmur32HashString/16
BenchmarkMurmur32HashString/16-8          	57295918	        21.48 ns/op	 744.91 MB/s	       4 B/op	       1 allocs/op
BenchmarkMurmur32HashString/64
BenchmarkMurmur32HashString/64-8          	25209924	        46.91 ns/op	1364.28 MB/s	       4 B/op	       1 allocs/op
BenchmarkMurmur32HashString/256
BenchmarkMurmur32HashString/256-8         	 7994089	       150.1 ns/op	1705.01 MB/s	       4 B/op	       1 allocs/op
BenchmarkMurmur32HashString/1024
BenchmarkMurmur32HashString/1024-8        	 2011392	       552.0 ns/op	1854.93 MB/s	       4 B/op	       1 allocs/op
BenchmarkMurmur128HashInt
BenchmarkMurmur128HashInt-8               	23675829	        51.75 ns/op	      40 B/op	       2 allocs/op
BenchmarkMurmur128HashInt32
BenchmarkMurmur128HashInt32-8             	24676378	        48.81 ns/op	      40 B/op	       2 allocs/op
BenchmarkMurmur128HashInt64
BenchmarkMurmur128HashInt64-8             	23552547	        52.26 ns/op	      40 B/op	       2 allocs/op
BenchmarkMurmur128HashBytes
BenchmarkMurmur128HashBytes/1
BenchmarkMurmur128HashBytes/1-8           	24905462	        47.82 ns/op	  20.91 MB/s	      40 B/op	       2 allocs/op
BenchmarkMurmur128HashBytes/4
BenchmarkMurmur128HashBytes/4-8           	25586134	        48.27 ns/op	  82.86 MB/s	      40 B/op	       2 allocs/op
BenchmarkMurmur128HashBytes/16
BenchmarkMurmur128HashBytes/16-8          	23899525	        49.53 ns/op	 323.03 MB/s	      40 B/op	       2 allocs/op
BenchmarkMurmur128HashBytes/64
BenchmarkMurmur128HashBytes/64-8          	21288523	        56.11 ns/op	1140.58 MB/s	      40 B/op	       2 allocs/op
BenchmarkMurmur128HashBytes/256
BenchmarkMurmur128HashBytes/256-8         	15894333	        78.09 ns/op	3278.34 MB/s	      40 B/op	       2 allocs/op
BenchmarkMurmur128HashBytes/1024
BenchmarkMurmur128HashBytes/1024-8        	 6944110	       175.5 ns/op	5836.18 MB/s	      40 B/op	       2 allocs/op
BenchmarkMurmur128HashString
BenchmarkMurmur128HashString/1
BenchmarkMurmur128HashString/1-8          	24699285	        49.38 ns/op	  20.25 MB/s	      40 B/op	       2 allocs/op
BenchmarkMurmur128HashString/4
BenchmarkMurmur128HashString/4-8          	23062286	        51.86 ns/op	  77.13 MB/s	      40 B/op	       2 allocs/op
BenchmarkMurmur128HashString/16
BenchmarkMurmur128HashString/16-8         	21514398	        52.96 ns/op	 302.09 MB/s	      40 B/op	       2 allocs/op
BenchmarkMurmur128HashString/64
BenchmarkMurmur128HashString/64-8         	18408691	        63.85 ns/op	1002.42 MB/s	      40 B/op	       2 allocs/op
BenchmarkMurmur128HashString/256
BenchmarkMurmur128HashString/256-8        	10089026	       128.9 ns/op	1985.29 MB/s	      40 B/op	       2 allocs/op
BenchmarkMurmur128HashString/1024
BenchmarkMurmur128HashString/1024-8       	 3612316	       302.6 ns/op	3383.88 MB/s	      40 B/op	       2 allocs/op
```

## License

This project is under the MIT license. See the [LICENSE](LICENSE) file for details.
