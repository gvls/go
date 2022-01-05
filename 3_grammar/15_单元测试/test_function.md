## 2 test function
### 3  format
First alphabet of `Name` must `大写` 
Follow must write in file which name `<name>_test.go` 
Follow must write in `package` which `function` which we test   stay

```go
func Test<Name>(t *testing.T){
	...
	t.<method>
	...
}

// t.Fatalf()				false   输出日志，并退出
// t.Logf(<output message>)		success 输出日志
```


### 3  test
It test all `Test<Name>` and if all is pass , result is `PASS` otherwise is `FAIL`
`-v` : `不管正确还是错误，都输出日志` 
not add `-v` : `输出错误日志` 

* test all
```shell
go test [-v]
```

* test one file
```go
go test [-v] <name>_test.go <name>.go
```

* test one `method` 
```go
go test [-v] -test.run <method>
```
