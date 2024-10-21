# matrixcpucache

TODO: find explanation for the differences between benchmarks:

- why are "cache miss" and "contention" scenarios more penalized on linux
- why does "false sharing" scenario not apply on mac

## Benchmark

Example on a MacBook Pro M1 Max:

```
Benchmark_sumLinked-10          	    3820	    300686 ns/op	       0 B/op	       0 allocs/op
Benchmark_sumCacheMiss-10       	    4084	    310187 ns/op	       0 B/op	       0 allocs/op
Benchmark_sumCacheHit-10        	    6340	    187508 ns/op	       0 B/op	       0 allocs/op
Benchmark_sumContention-10      	    6734	    185997 ns/op	     987 B/op	      22 allocs/op
Benchmark_sumFalseSharing-10    	   13892	     87357 ns/op	     816 B/op	      22 allocs/op
Benchmark_sumParallel-10        	   14192	     84443 ns/op	     936 B/op	      23 allocs/op
```

Example on a Linux AMD EPYC 4344P with DDR5 3600MHz:

```
Benchmark_sumLinked-16          	    5544	    217824 ns/op	       0 B/op	       0 allocs/op
Benchmark_sumCacheMiss-16       	    1418	    802728 ns/op	       0 B/op	       0 allocs/op
Benchmark_sumCacheHit-16        	   10000	    112634 ns/op	       0 B/op	       0 allocs/op
Benchmark_sumContention-16      	    3865	    267992 ns/op	    1576 B/op	      34 allocs/op
Benchmark_sumFalseSharing-16    	    8581	    136781 ns/op	    1297 B/op	      34 allocs/op
Benchmark_sumParallel-16        	   19845	     59506 ns/op	    1416 B/op	      35 allocs/op
```

## Perf stat

Example on a Linux AMD EPYC 4344P with DDR5 3600MHz:

```
 Performance counter stats for './matrixcpucache.test -test.run=XXX -test.bench=Benchmark_sumLinked -test.benchmem':

        1462167548      cache-references
           8237463      cache-misses                     #    0.563 % of all cache refs
        6670806625      cycles
        8012768076      instructions                     #    1.20  insn per cycle

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

 Performance counter stats for './matrixcpucache.test -test.run=XXX -test.bench=Benchmark_sumContention -test.benchmem':

         600993441      cache-references
         102095713      cache-misses                     #   16.988 % of all cache refs
       55611233599      cycles
       22127794351      instructions                     #    0.40  insn per cycle

 Performance counter stats for './matrixcpucache.test -test.run=XXX -test.bench=Benchmark_sumFalseSharing -test.benchmem':

        1050209011      cache-references
         107320423      cache-misses                     #   10.219 % of all cache refs
       43731622987      cycles
       37461357752      instructions                     #    0.86  insn per cycle

 Performance counter stats for './matrixcpucache.test -test.run=XXX -test.bench=Benchmark_sumParallel -test.benchmem':

        3148097291      cache-references
          88023502      cache-misses                     #    2.796 % of all cache refs
       25877437876      cycles
      133755133003      instructions                     #    5.17  insn per cycle
```
