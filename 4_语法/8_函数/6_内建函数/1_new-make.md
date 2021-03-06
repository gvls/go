##  new 和 make
###   共同点
说明：在程序运行时，在堆中申请空间 赋值给变量。
性能： 申请太大空间会 影响程序的 性能，尽量一次申请代替 多次申请 来提高程序性能。

###   new
为 值类型/引用类型 申请空间。
返回值的类型 ： 申请数据类型的 指针类型
返回值       ： 申请数据类型的 **零值的指针**
```go
new(数据类型) *数据类型
```


###   make
只为 引用类型 申请空间 （`slice` `map` `channel` ）。
返回值的类型 ： 申请数据类型
返回值       ： 申请数据类型的 **非零值**
```go
make(数据类型, 参数) 数据类型
```

###   典型例子
* `a` 为`*[]int` 零值的指针，不能`append` 。
```go
a := new([]int)
a = append(a, 9)
```

