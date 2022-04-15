##  recover 
捕获当前位置到之间 并在函数内的 错误
处理抛出的错误，并让程序继续执行下去

###   格式
```go
err := recover()
if err != nil {
	// 处理错误
}
```
```go
defer func () {
	// 捕获错误
}()
```
