# Performance Cheatsheet

## Benchmarking

Benchmarking is helpful when measuring, optimizing or debugging the performance of a function. Benchmarks reveal information about which parts of code are slow or require a lot of memory.

### Optimization Process

**1. Is this code/app slow**

Always ask yourself this as a first step, there's no benefit to optimizing something that's fast enough or cannot be made significantly faster. Performance optimizations can come with added complexity, making code harder to understand or maintain, most of the time this is not a useful tradeoff. 

What does fast mean for the context you're working in? How important is speed vs. understandability?

**2. Find performance bottlenecks**

Optimizing performance is all about addressing bottlenecks. It's easy to see code that _looks_ slow & jump to the conclusion that it is. Putting time & energy into optiziming the wrong area can lead to frustration & wasted effort.

The best way to find bottlenecks is by profiling your code. Often cpu profiling (how much time the cpu spends in different spots of the code) is a good place to start, moving to memory profiling if you see a lot of references to `malloc` methods.  Running the tracer can be useful for apps that rely heavily on go routines.

**3. Reliably measure current performance**

Profile or run benchmarks on the existing code.  See more on running reliable benchmarks below.

**4. Make sure the code is unit tested**

Faster code that doesn't work properly defeats the purpose of performance optimization.

**5. Iteratively Optimize**

- Profile/Benchmark to find bottleneck
- Optimize
- Profile/Benchmark again

## Benchmarks

**Example**

```
func BenchmarkMyFunc(b *testing.T) {
  // report memory allocations in output
  b.ReportAllocs()

  for n := 0;n < b.N; n++ {
    myFunc()
  }
}
```
Benchmarks are run by passing the `-bench` flag when running `go test`, for example `go test -run $$$ -bench BenchmarkMyFunc` would run this benchmark. Note the `-run $$$` here which excludes running unit tests (since no tests start with `$$$`).

The benchmark function must run the target code b.N times. During benchmark execution, b.N is adjusted until the benchmark function lasts long enough to be timed reliably.

### Iterating with Benchmarks

When running benchmarks multiple times they will tend to get slower because **__laptops throttle__**. This slow-down has nothing to do with the code itself & can result in false positives or negatives. Using [benchstat](https://godoc.org/golang.org/x/perf/cmd/benchstat) and running benchmarks multiple times can help ensure we are accurately measuing the speed of the code (not the machine).

With go test we can run series of Benchmarks consecutively with the `-count` flag which we can then save to a file for future comparison. This will output how long the function takes to execute and the number of memory allocations on average.
```
// Run benchmark multiple times & output results to file
go test -run $$$ -bench <benchmarkFunc> -count=10 > <filename>
go test -run $$$ -bench BenchmarkMyFunc -count=10 > beforeChanges
```

see [slower](https://github.com/msiehoff/gopher-notes/blob/perf-profile/performance/slower) for example output.

Benchstat shows us variance in groups of runs. The higher the variance, the less reliable the results.

```
// Use benchstat to see the variation in the benchmark results
benchstat <filename>
```

```
// Optimize then rerun benchmarks
go test -run $$$ -bench BenchmarkMyFunc -count=10 > afterChanges

// Compare benchmark output files with benchstat
benchstat beforeChanges afterChanges
```

![image](https://user-images.githubusercontent.com/901644/66795025-f43cb580-eec8-11e9-898e-b804babaaba3.png)

### Profiling with Benchmarks

```
// run benchmark & save cpu profile
go test -run $$$ -benchmark BenchmarkMyFunc -cpuprofile <filename>

// run benchmark & save memory profile
go test -run $$$ -benchmark BenchmarkMyFunc -memprofile <filename>

// view profile in browser
go tool pprof -http=":<port>" <filename>"
```

## Profiling

### API's

### Other Programs
