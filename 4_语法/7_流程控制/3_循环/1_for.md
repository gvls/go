##  for 
可以遍历 `string` 、`array` 、 `slice` 、 `channel` 

* 语句A
只执行一次
在第一次 进入前 执行

* 条件
在每次 进入前 执行
可以用返回值类型是`bool` 的语句代替

* 语句B
在每次循环体 执行完 后执行



###   格式
```go
for [语句A] ; [条件] ; [语句B] {
	...
}
```
```go
[语句A]
for <条件> {		// 像 while
	...	
	...
	...
	[语句B]
}
```
```go
[语句A]
for {				// for {   等同   for ; ; {
	...
	...
	...
	if <条件> { 	// 像 do-while
		break
	}
}
```
