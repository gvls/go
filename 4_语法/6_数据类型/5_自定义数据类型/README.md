##  README
定义一种**全新**的数据类型，新数据类型 ！= 旧数据类型
在全局里 定义
如果新数据类型的名字 首字母大写，该新数据类型 可以被其他包访问

```go
type 新数据类型 某数据类型
```

###   例子
```go
type myInt int
var a myInt
var b int
a = b 			// 抛出异常
```
