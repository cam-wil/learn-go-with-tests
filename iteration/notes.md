1. the only loops that exist in go are for loops, but they can be setup as while loops
2. `for i := 0; i < SOME_NUM; i++ {...}`
3. Benchmarks. Start function like `BenchmarkXxx(b *testing.B)`
4. `b.N` is an automated iterator for the benchmark
5. `go test -bench=.` to run benchmarks