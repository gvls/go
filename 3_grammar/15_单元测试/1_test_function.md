##  test function
###   format
First alphabet of `Name` must `大写` 
Name of file **must** is `<funcName>_test.go` 
`package` **must** is `package` which `function` which we test

```go
func Test<funcName>(t *testing.T){
	...
	t.<method>
	...
}

// t.Fatalf()					false   输出日志，并退出
// t.Errorf()
// t.Logf(<output message>)		success 输出日志
```


###   test
It test all `Test<funcName>` and if all is pass, result is `PASS` otherwise is `FAIL`
not add `-v` : `输出错误日志` 
`-v` : `不管正确还是错误，都输出日志` 

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

* get test cover
```go
go test -cover [-coverprofile=<fileName>]
// go tool cover -html=<fileName> 			// get visualization of cover
```



###   son test
```go
t.Run(xxx, func(t *testing.T){
	....
})
```


