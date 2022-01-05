## 2 memory
Attention : give `variate` to `function` and `function` add or remove element or do something, old `slice` **not change**
Like :
```go
type slice struct {
	ptrName unsafe.Pointer
	lenName int
	capName int
}
```

### 3  `address` of first element 
Pointer point to first element of `arrary` 

### 3  length of `slice` all element
number of element of   `array` which `slice` quote to

### 3  capacity of `slice`
capacity of   `arrary` which `slice` point to

### 3  `function参数传递slice` 
`funciton里面的形参slice复制了实参slice的内容（数组的地址，长度，容量），改变形参slice会改变实参slice` 
