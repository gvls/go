## 2 `interface`
~~It is `pointer`~~
`实现了接口的数据类型` like `interface` 
Only can define `data type` and Can't create `variate` directly
Only can bind to `实现了这个接口的数据类型的实例` 
Can't define `variate` in `interface` 

### 3  `interface` 
One `interface` can inherit other `interface` more than one. If want to implement current `interface` , must implement father `interface` of current `interface` 

1. Have some `method` which isn't implement (`定义一种统一的标准规范给大家使用`)
Like `USB` 
```go
interface {}	// this data type of interface can be all data type implement
```

```go
type interfaceName interface{			// interfaceName is a new data type
	[data type]				// inherit other interface
	methodNameA(paramterList) returnList	// 没有实现的方法
	...
}
```



### 3  `实现了接口的数据类型` (`不仅仅是struct`)
1. `data type` **have** method same as all method of `interface` (`遵循了interface的标准，就 可以是 interface的实现数据类型`)
Like `键盘` , `鼠标` 

```go
type intName int
type structName struct {
	...
}
func (typeVariateName typeName) methodNameA(parameterList) returnList{
	...
}
```




### 3  `使用接口` 
When use `interface` to receive value of `实现了接口的数据类型` , `variate` of `interface` only can use `method` of `interface` and `method` of it's father `interface` 
1. `variate` of `interface` be evaluation by value of `实现了接口的数据类型` 
2. use `variate` to use `method` of `interface` 
Like `电脑调用USB` (`USB 连接的是鼠标还是键盘，由 实现了接口的数据类型决定`)

```go
var variateName interfaceName = `实现了接口的数据类型的变量的值` 
variateName.methodNameA(parameterList)	// 执行 实现了接口的数据类型 的自定义的操作
```



### 3  tips
one custom `daty type` can implement more than one `interface` 
`interface` can as `形参` in `function` to receive any `data type`  
In old Go version, When `interface` inherit `interface` , can't appear more than one `method` same 
Attention `是` `dataType` `还是` `*dataType` implement `method` of `interface` 
```go
type A interface{
	test()
}
type Stu struct {
}
func (s *Stru) test(){
	fmt.Println("test")
}
func main() {
	var s Stu 
	var a A = s	// `s` must be replace by `&s` 
	s.test()
}
```



### 3  function
`设计一种规范（方法），让大家去遵守` 
