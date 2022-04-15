##  main
程序的入口
main函数所在的包 必须为main包

###   格式
```go
func main(){

}
```

###   简单参数
需要导入 `os` 包
参数的顺序是 有序的
```go
os.Args		// 类型是 []string，第一个是 二进制文件的名字
```


###   复杂参数
需要导入 `flag` 包
参数的顺序 可以随意
每一个参数前需要添加`-arg` 去标识

* 获取`string` 类型
```go
flag.StringVar(&变量, "arg", 默认值, "参数的说明") // 默认值 ： 若获取不到参数，就使用默认值 赋给变量
flag.Parse()	// 从 os.Args[1:] 获取参数
```
* 获取`int` 类型
```go
flag.IntVar(&变量, "arg", 默认值, "参数的说明")
flag.Parse()
```
