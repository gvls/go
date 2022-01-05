## 2 `method` 
Like `function` 
custom `data type` can have `method` 

* have effect in specific `data type` 
`method` bind specific `data type` 
If `method A` bind `data type A`, only `variate` of `data type A` can use `method A`

### 3  define
define out of all `function` 

#### 4   example `data type` 
```go
type dataTypeA struct {
	...
}
```

#### 4   `method` 
`partVariateName` often write `小写的` `dataTypeA`
`partVariateName` can not use
`dataTypeA` `这个数据类型有一个方法` `methodName`
`methodName` can repetition When `dataTypeA` is  different

* not change `property` of `variate` 
```go
// 不管调用这个methodName的variateNameA是variateNameA还是(&variateNameA)，partVariateName是值还是地址由它后面的类型决定

func (partVariateName dataTypeA) methodName(parameterList) returnList { 
	...
}
```

* can change `property` of `variate` 
```go
// 不管调用这个methodName的variateNameA是variateNameA还是(&variateNameA)，partVariateName是值还是地址由它后面的类型决定

func (partVariateName *dataTypeA) methodName(parameterList) returnList { 
	...
}
```



### 3  use
`variateNameA实参` give value to `partVariateName形参` and execute `methodName`

```go
// 不管method的partVariateName是不是指针类型都可以用这个方法来使用

var variteNameA dataTypeA
variateNameA.methodName(parameterList) 
```

```go
// 不管method的partVariateName是不是指针类型都可以用这个方法来使用

var variteNameA dataTypeA
(&variateNameA).methodName(parameterList) 
```



#### 4   `String()` 
If `data type` of `variate` implement this `method` , `fmt.Println(varaite)` will execute this `method` 
