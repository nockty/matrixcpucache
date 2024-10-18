# matrixcpucache

Example run:

```
Benchmark_sumLinked-10          	    3789	    305537 ns/op	       0 B/op	       0 allocs/op
Benchmark_sumCacheMiss-10       	    3591	    301356 ns/op	       0 B/op	       0 allocs/op
Benchmark_sumCacheHit-10        	    6360	    188417 ns/op	       0 B/op	       0 allocs/op
Benchmark_sumFalseSharing-10    	    6196	    183374 ns/op	     993 B/op	      22 allocs/op
Benchmark_sumParallel-10        	   15564	     82270 ns/op	     936 B/op	      23 allocs/op
```

