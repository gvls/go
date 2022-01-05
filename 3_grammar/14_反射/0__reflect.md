## 2 `reflect` 
When in running, `reflect` can get any information of `variate` 
Need `import "reflect"` 
`Kind` : `data type` belong which kind
`Type` : specific `data type` 

### 3  use scene
* `序列化struct` 
`自定义序列化后，属性的名字` 

* `函数适配器` 
`一个函数实现传入 任何函数变量和参数就可以执行` 

### 3  get `interface{}` by `varaite` 
evaluation



### 3  get type(`类型`) by `interface{}` 
`Type.Kind()` : `类别`(constant)

```go
reflect.TypeOf(variateName) reflect.Type
```


### 3  get `Value` by `interface{}` 
`variate` , `interface{}` , `reflect.Value` : they can `相互转换` 
`Value.Kind()` : `类别`(constant)
`parameter` of `refalect.ValueOf()` can is specific `data type` 

```go
reflect.ValueOf(interfaceVariateName interface{}) reflect.Value
```

* get number of `field`
```go
<Value>.NumField()
```

* get value of `field`
```go
<Value>.Field(<index>)
```

* get tag of `field`
```go
<Type>.Field(<index>).Tag.Get(<key>)
```

* get number of `method` 
```go
<Value>.NumMethod()
```

* get `method` by index
```go
<Value>.Method(index int) Value		// index is define after order name of method
```

* use `method` 
```go
<Method>.Call(in []Value) []Value
```





### 3  get `Value` by `Pointer Value` or `interface Value` 
When use `SetString()` , must pass `Pointer Value` to `Value` 
```go
reflect.Elem() Value
```




### 3  get `interface{}` by `Value` 
```go
reflectVariateName.Interface()	interface{}
```


### 3  get `variate` by `interface{}` 
`断言` 


