##  add element
###   append 
basic on **make** 

* add value
```go
append(slice1, newElement [, ...]) 
```

* add slice 
```go
append(slice1, newSlice...) 	  
```

####    If cap1 - len1   `>=`   count of newElement 
1. Add newElement to end of arrary1
2. function return :
> arrary   : **array1**
> length   : len1 + number of newElement
> capacity : cap1 + number of newElement

####    If cap1 - len1   `<`    count of newElement 
If length of array <  1024, append() will new 2xlen1 		array
If length of array >= 1024, append() will new 1.25xlen1 array

1. Create **new** 匿名arrary 
2. This new 匿名数组 copy all element from arrary1
3. Add newElement to end of this new 匿名数组
4. function return :
> arrary   : this **new** 匿名数组
> length   : len1 + number of newElement
> capacity : cap1 + number of newElement


####    tips
* 如果追加元素给了slice1，尽量避免再追加元素给slice1
If cap1 - len1 >= number of newElement.
slice2 := append(slice1, newElement2)
slice3 := append(slice1, newElement3) 
array2 == array3 and newElement2 will be change to newElement3
So 尽量把新追加后的value赋值给原来的slice 

* 若slice是切现有的一个数组，尽量不要追加新元素 

* 函数传递slice 过去后，函数append不会影响原来的slice
