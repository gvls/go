## 2 `struct`
It be use to create a `data type` 
It can include any `dataType` and number

* **支持**  `面向对象编程OOP` 
`struct` be use as `类` 
`variate` of `struct` be use as `对象` | `实例` 

### 3  define custom `struct` 
define in `globle` 

`structName大写首字母` : other `package` can visit
`variateName大写首字母` : other `package` can visit
```go
type <sturtName> struct {
	<variateName> <variateType>
	[variateName, variateName variateType]
	...
	<variateNamex> <variateTypex>
}
```


### 3  define `variate` of `struct` 
```go
var variateName structType
```


### 3  define and initialize `variate` 

* initialize 
```go
var variateName = structName{}
```

* initialize by value
must write value of all `property` 
```go
var variateName = structName{value1, value2 ...}
```

* initialize by value which point property
can write value of part `property` 
```go
var variateName = structName{
	property1 : value1 ,
	property2 : value2 ,
	...
}
```

* initialize by ~~nil~~ and get pointer
```go
var variateName *structName = &structName{}
// use is : (*variateName).property   or   variateName.property(编译器会转换)
```

* get pointer
```go
var variateName *structName = new(structName)
// use is : (*variateName).property   or   variateName.property(编译器会转换)
```




### 3  use
If `data type` of `structName` is `pointer` , we also can use follow method (go will handle it)


* visit all property
If `structName` is `pointer` , follow output : `&all element` 
```go
structName
```

* visit `property` or evaluation
If `structName` is `pointer` of `struct`, we also can use follow way to visit `property` (`go` will handle it in `build`)
```go
structName.propertyName
```



### 3  `struct A` => `struct B` 
`b = B(a)` 
`property` of A must same as `property` of B (name, number, type)

