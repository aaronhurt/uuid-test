# uuid-test


Quick UUID package performance testing. Look at `bench_test.go` for packages and test cases. The formatter and included libs are licensed the same as the original sources.

## Results

```
$ go test -bench . -benchtime 10s
BenchmarkV4Google-4            	10000000	      1287 ns/op
BenchmarkV4Satori-4            	10000000	      1273 ns/op
BenchmarkV4m4rw3r-4            	10000000	      1295 ns/op
BenchmarkFastUUIDRaw-4         	1000000000	        18.9 ns/op
BenchmarkFastUUIDFormatted-4   	200000000	        92.5 ns/op
PASS
ok  	github.com/leprechau/uuid-test	91.061s
```
