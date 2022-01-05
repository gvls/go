## 2 `slice`
It is a `struct` and `quote type` and basic on `array` 
`slice` can quote more than one `array` 
Use it is like `array` 
It's **length can be automatic and dynamic change**(when capacity > length) in procedure running
capacity >= length



### 3  define
```go
var sliceName []dataType	
```



### 3  evaluation or initialize
* in coding point to exist `array` 
`slice` point to `已经存在的数组` and initialize slice by length and capacity
`切一个数组得到一个切片` 
`切片可以继续被切片` 
```go
// if indexi == 0 , indexi can not write
// if indexj == len(arrary) , indexj can not write
sliceName = arraryName[indexi:indexj]	// [indexi,indexj)
```

* in coding point to create `array` 
`slice` point to `现在创建的匿名的数组` and initialize slice by length and capacity
```go
sliceName = []dataType {value1, value2 ...}
```

* point value of specific element by `index` 
```go
sliceName = []dataType {
	index : value ,
	index : value ,
	...
}
```

* in running to new `array` 
`slice` point to `程序运行时新建并分配的一个底层的匿名的数组` and initialize slice by length and capacity
```go
sliceName = make([]dataType, len [, cap])
```

### 3  define and evaluation

* in add element
```go
var sliceName []dataType
sliceName = append(sliceName, newElement)
```


