# fat contexts

This repo holds code for [Context-induced performance bottleneck in Go](https://gabnotes.org/fat-contexts/).

Reproduce my results:

```console
$ go test -bench=.
goos: darwin
goarch: arm64
pkg: github.com/Crocmagnon/fat-contexts
BenchmarkContext/shadow_10-8         	 2261456	       763.0 ns/op
BenchmarkContext/fat_10-8            	 1662235	       743.6 ns/op
BenchmarkContext/shadow_100-8        	  236544	      4888 ns/op
BenchmarkContext/fat_100-8           	   53778	     22275 ns/op
BenchmarkContext/shadow_1000-8       	   24499	     48474 ns/op
BenchmarkContext/fat_1000-8          	     711	   1698109 ns/op
BenchmarkContext/shadow_10000-8      	    2472	    489804 ns/op
BenchmarkContext/fat_10000-8         	       6	 170819118 ns/op
BenchmarkContext/shadow_100000-8     	     248	   4938549 ns/op
BenchmarkContext/fat_100000-8        	       1	17150788208 ns/op
PASS
ok  	github.com/Crocmagnon/fat-contexts	31.454s
$ go run ./generate/main.go > data
$ go run ./plot/main.go
$ open plot.html
```

Sorry there are a lot of hardcoded file names in here, but since I don't intend on reusing this code I won't fix it.
