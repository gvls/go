##  test performance
Call many count of `function` and calculate **average** execute time or memory use

###   format
First alphabet of `Name` must `大写` 
Follow must write in file which name `<funcName>_test.go` 
Follow must write in `package` which `function` which we test   stay

```go
func Benchmark<funcName>(b *testing.B){
	...
	b.<method>
	...
}

// t.ResetTimer()			 reset time
```



###   test
####    test specific
```go
go test -bench=<funcName> [option]
```

* option
`-count <n>` : count
`-cpu <n>` : format of `cpu`
`-benchmem` : used of memory

####    test all
```shell
go test -bench=. [option]
```

####    out put
```shell
Benchmark<funcName>-<physicsThread> <count of execute>	<time of once execute> [size of memory once] [apply of memory once]
```
