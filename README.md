# fat contexts

This repo holds code for [On context-induced performance bottleneck in Go](https://gabnotes.org/fat-contexts/).

Reproduce my results:

```console
$ go test -bench=.
goos: darwin
goarch: arm64
pkg: trash
BenchmarkContext/shadow_1000-8         	   19352	     61271 ns/op
BenchmarkContext/fat_1000-8            	     532	   2187010 ns/op
BenchmarkContext/shadow_10000-8        	    1903	    639371 ns/op
BenchmarkContext/fat_10000-8           	       5	 219400100 ns/op
BenchmarkContext/shadow_100000-8       	     194	   6374344 ns/op
BenchmarkContext/fat_100000-8          	       1	21851940167 ns/op
PASS
ok  	trash	30.801s
$ go run ./generate/main.go > data
$ go run ./plot/main.go
$ open plot.html
```

Sorry there are a lot of hardcoded file names in here, but since I don't intend on reusing this code I won't fix it.
