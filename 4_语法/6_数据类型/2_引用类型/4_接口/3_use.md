##  use
When use variate of `interface` to receive value of `实现了接口的数据类型`, variate of `interface` **only can use** `method` of `interface` and `method` of it's father `interface` 
Like `电脑调用USB` (`USB 连接的是鼠标还是键盘，由 实现了接口(method)的数据类型决定`)


###   process
1. variate of `interface` be evaluation by value of `实现了接口的数据类型` 
2. use variate to use `method` of `interface` 

```go
var variateName interfaceName = `实现了接口的数据类型的变量的值` 
variateName.methodNameA(parameterList)	// 执行 实现了接口的数据类型 的自定义的操作
```
