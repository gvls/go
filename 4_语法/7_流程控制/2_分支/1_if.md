##  if 
可以嵌套使用

* 条件
推荐不使用`()` 包裹
可以在条件前面添加 多个`;`结尾的语句
条件的结果一定要是`bool`类型
不能在这里定义变量



###   单分支 
推荐使用的方式，经常用来处理错误的情况
```go
if 条件 {
	...
}
```



###   双分支
```go
if 条件 {
	...
} else {
	...
}
```



###   多分支 
```go
if 条件 {
	...
} else if 条件 {
	...
}...
...
} else {
	...
}
```
