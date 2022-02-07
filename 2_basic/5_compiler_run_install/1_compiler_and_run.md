##  compiler and run
It generate a **binary** file which add dependences library of runtime in current directory
We can copy it to other machine to run and don't need install environment of `go develop`

###   process
Source code -> `go build` -> binary file -> `run` -> get result

###   compiler
If have error on compiler, it show line of error
`Go` will find `<name>.go` from `$GOPATH/src/` 

* compiler one file
```go
[parameter] go build [-o new name] <name>.go			//<name>.go need include `main function`
```
* compiler project
```go
[parameter] go build [-o [path]<new name>] <path>/main	//main directory need include `pacakge main`
```
* compiler `./` 
```go
[parameter] go build 									// parameter is : GOOS=xxx ...
```
* compiler small
`ags "-s -w"参数去掉符号表和调试信息，一般能减小20%的大小。`

###   run
```shell
./<name>
```
