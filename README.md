# matrixcpucache

## Benchmark

Example on a MacBook Pro M1 Max:

```
Benchmark_sumLinked-10          	    3789	    305537 ns/op	       0 B/op	       0 allocs/op
Benchmark_sumCacheMiss-10       	    3591	    301356 ns/op	       0 B/op	       0 allocs/op
Benchmark_sumCacheHit-10        	    6360	    188417 ns/op	       0 B/op	       0 allocs/op
Benchmark_sumFalseSharing-10    	    6196	    183374 ns/op	     993 B/op	      22 allocs/op
Benchmark_sumParallel-10        	   15564	     82270 ns/op	     936 B/op	      23 allocs/op
```

Example on a Linux AMD EPYC 4344P with DDR5 3600MHz:

```
Benchmark_sumLinked-16          	    5557	    216467 ns/op	       0 B/op	       0 allocs/op
Benchmark_sumCacheMiss-16       	    1399	    855129 ns/op	       0 B/op	       0 allocs/op
Benchmark_sumCacheHit-16        	   10000	    111818 ns/op	       0 B/op	       0 allocs/op
Benchmark_sumFalseSharing-16    	    3813	    269957 ns/op	    1580 B/op	      34 allocs/op
Benchmark_sumParallel-16        	   20118	     59674 ns/op	    1416 B/op	      35 allocs/op
```

## Perf stat

Example on a Linux AMD EPYC 4344P with DDR5 3600MHz:

```
 Performance counter stats for './matrixcpucache.test -test.run=XXX -test.bench=Benchmark_sumCacheMiss -test.benchmem':

         717608656      cache-references
         133408844      cache-misses                     #   18.591 % of all cache refs
        8038637993      cycles
        6539285108      instructions                     #    0.81  insn per cycle

 Performance counter stats for './matrixcpucache.test -test.run=XXX -test.bench=Benchmark_sumCacheHit -test.benchmem':

         849241795      cache-references
           1677349      cache-misses                     #    0.198 % of all cache refs
        5999537254      cycles
       37135031090      instructions                     #    6.19  insn per cycle

 Performance counter stats for './matrixcpucache.test -test.run=XXX -test.bench=Benchmark_sumFalseSharing -test.benchmem':

         600993441      cache-references
         102095713      cache-misses                     #   16.988 % of all cache refs
       55611233599      cycles
       22127794351      instructions                     #    0.40  insn per cycle

 Performance counter stats for './matrixcpucache.test -test.run=XXX -test.bench=Benchmark_sumParallel -test.benchmem':

        3148097291      cache-references
          88023502      cache-misses                     #    2.796 % of all cache refs
       25877437876      cycles
      133755133003      instructions                     #    5.17  insn per cycle

```
