# Performance Cheatsheet

## Benchmarking

Benchmarking is helpful when measuring, optimizing or debugging the performance of a function. Benchmarks reveal information about which parts of code are slow or require a lot of memory.

### Optimization Process

*1. Is this code/app slow?*
Always ask yourself this as a first step, there's no benefit to optimizing something that's fast enough or cannot be made significantly faster. Performance optimizations can come with added complexity, making code harder to understand or maintain, most of the time this is not a useful tradeoff. 

What does fast mean for the context you're working in? How important is speed vs. understandability?

*1. Find performance bottlenecks*
Optimizing performance is all about addressing bottlenecks. It's easy to see code that _looks_ slow & jump to the conclusion that it is. Putting time & energy into optiziming the wrong area can lead to frustration & wasted effort.

The best way to find bottlenecks is by profiling your code. Often cpu profiling (how much time the cpu spends in different spots of the code) is a good place to start, moving to memory profiling if you see a lot of references to `malloc` methods.  Running the tracer can be useful for apps that rely heavily on go routines.

*1. Get a reliable starting place*


*1. Make sure the code is unit tested*

Faster code that doesn't work properly defeats the purpose of performance optimization.

*1. Iteratively Optimize*

### Example Benchmark

```
func BenchmarkMyFunc(b *testing.T) {
  // report memory allocations in output
  b.ReportAllocs()

  for n := 0;n < b.N; n++ {
    myFunc()
  }
}
```
Benchmarks are run by passing the `-bench` flag when running `go test`, for example `go test -run $$$ -bench BenchmarkMyFunc` would run this benchmark.

The benchmark function must run the target code b.N times. During benchmark execution, b.N is adjusted until the benchmark function lasts long enough to be timed reliably.

### Useful Benchmark Commands

```
// Save binary for future comparison

// Run benchmark multiple times & output results to file

// Run benchmark using benchstat to get variance

// Compare benchmark output files with benchstat
```

## Profiling

```
// run benchmark & save cpu profile
go test -run $$$ -benchmark BenchmarkMyFunc -cpuprofile <filename>

// run benchmark & save memory profile
go test -run $$$ -benchmark BenchmarkMyFunc -memprofile <filename>

// view profile in browser
go tool pprof -http=":<port>" <filename>"
```

