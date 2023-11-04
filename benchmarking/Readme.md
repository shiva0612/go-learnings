## benchmarking
```
go test -bench=. -cpu=1 -count=3 -benchtime=2s -cpuprofile=cpu.out -memprofile=mem.out -benchmem
    -cpu=1,2,3
    -benchtime=2s,100x
    -bench=BenchmarkFunc1 pkgName

go tool pprof memprofile.out
top
list function_name

brew install graphviz
go tool pprof -http=:8080 cpu.out
go tool pprof -http=:8080 mem.out
--------------------------------------------------
go test -bench=. -benchmem > old.txt
go test -bench=. -benchmem > new.txt
benchstat old.txt new.txt
//you can also benchstat for cpuProfile as well 

```
