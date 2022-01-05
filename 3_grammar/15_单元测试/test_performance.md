## 2 test performance
### 3  format
First alphabet of `Name` must `大写` 
Follow must write in file which name `<name>_test.go` 
Follow must write in `package` which `function` which we test   stay

```go
func Benchmark<Name>(b *testing.B){
	...
	b.<method>
	...
}

// t.ResetTimer()			 reset time
```



### 3  test

#### 4   test all
```shell
go test -bench=. [option]
```

* option
`-count <n>` : count
`-cpu <n>` : format of `cpu`
