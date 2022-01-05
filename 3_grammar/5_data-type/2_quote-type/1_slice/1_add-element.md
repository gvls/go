## 2 add element
### 3  `append` 
basic on `make` 
`假设: slice1 : arrary is arrary1.   length is len1.   capacity is cap1`

* add value
```go
append(slice1, newElement [, ...]) slice2
```

* add `slice` 
```go
append(slice1, newSlice...) slice2
```

#### 4   If `cap1` - `len1`   `大于或等于`   number of `newElement` 
1. Add `newElement` to end of `arrary1`
2. `function` return :
> `arrary` : `array1` 
> `length` : `len1` + number of `newElement` 
> `capacity` : `cap1` + number of `newElement` 

#### 4   If `cap1` - `len1`   `小于`   number of `newElement` 
1. Create **new** `匿名arrary` 
2. This new `匿名数组` copy all element from `arrary1`
3. Add `newElement` to end of this new `匿名数组`
4. `function` return :
> `arrary` : this new `匿名数组` 
> `length` : `len1` + number of `newElement` 
> `capacity` : `cap1` + number of `newElement` 

#### 4   tips
* `如果追加元素给了slice1，尽量避免再追加元素给slice1` because :`If cap1 - len1 >= number of newElement slice2 := append(slice1, newElement2) slice3 := append(slice2, newElement3) ,  array2 == array3 and newElement2 will be change to newElement3` so `尽量把新追加后的value赋值给原来的slice` 

* `若slice是切现有的一个数组，尽量不要追加新元素` 

