##  value
we **must evaluation** before use variate

###   空slice
```go
[]dataType{}
```
```go
make([]int, 0)
```
```shell
physics :
   		_________________________
		|       |       |       |
slice	|pointer|   2   |   4   |
		|       |       |       |
		-------------------------
		    |
			v
		  xxxxx
```

###   point to **exist** array or slice
slice point to 已经存在的数组 and initialize slice by length and capacity
切一个数组得到一个切片 
切片可以继续被切片 
if not write indexi, default value is 0
if not write indexj, default value is len(array)
indexj can't  >  length of array or cap of slice
length = indexj - indexi
new array == old array

* default cap
```go
arraryOrSlice[indexi:indexj]	// [indexi,indexj) cap default is cap(arrayOrSlice)-indexi
```
* point cap
```go
arraryOrSlice[indexi:indexj:cap]	// [indexi,indexj)
```

###   point to **create** array 
slice point to 现在创建的匿名的数组 and initialize slice by length and capacity
```go
[]dataType{value1, value2 ...}
```

###   point value of specific element by index 
```go
sliceName = []dataType {
	index : value ,
	index : value ,
	...
}
```

###   pointer
add `&` to front of pre way

###   in running to new array 
slice point to `程序运行时新建并分配的一个底层的匿名的数组` and initialize slice by length and capacity
```go
make([]dataType, len [, cap])	// if not write cap, cap == len
```

###   make and add element
```go
var sliceName []dataType
sliceName = append(sliceName, newElement)
```


