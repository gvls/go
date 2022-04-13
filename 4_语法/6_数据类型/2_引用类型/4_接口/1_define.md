##  define
Like `USB` 

* Have some `method` which isn't implement
定义一种统一的标准规范给大家使用

* One `interface` can inherit other `interface` more than one.
If want to implement current `interface` , must implement father `interface` of current `interface` 
In old Go version, When `interface` inherit `interface`, can't appear more than one `method` same 

###   have method
```go
type interfaceName interface{				// interfaceName is a new data type
	[data type]								// inherit other interface
	methodNameA(paramterList) returnList	// 没有实现的方法
	...
}
```

###   interface{}
this can as `形参` in `function` to receive any `data type`  
```go
interface {}	// this data type of interface can be all data type implement
```
