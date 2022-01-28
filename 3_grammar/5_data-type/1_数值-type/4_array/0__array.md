## 2 `array`
```shell
1  9  3 
_____________________________
|   |   |   |   |   |   |   |
| 1 | 9 | 3 | 0 | 0 | 0 | 0 |
|   |   |   |   |   |   |   |
-----------------------------
```
It can save many value which is same `data type` 
When define, Length of it is fixation
`address` of `array` is `address` of first element of `array` 


### 3  `data type` format
`dataType` can is any `data type` 
Because of `dataType` is different, go have many `array data type` 
`不同长度的数据类型也是 不同的数据类型` 
`index of array` is from 0 ~ length-1
```go
[length of arrary]dataType
```


### 3  define
```go
var arrayName [length of array]dataType	
```


### 3  evaluation
`length of array` must same as length of `array variate`
`array` is `数值类型` , so `array` is use like `int` ...

* not write `data type` of `variate` 
```go
arrayName = [length of array]dataType {value, value ...}	// 读起来就像： 多少个 dataType 值分别是
```

* not point `length of array` 
```go
arrayName = [...]dataType {value, value ...}
```

* point value of specific element by `index` 
```go
arrayName = [length of array]dataType {
	index : value ,
	index : value ,
	...
}
```

* get `pointer` and evaluation in running
```go
arrayNmae = new([length of arry]dataType)
```



### 3  define and    initialize or evaluation


### 3  use

* visit all element
```go
arrayName
```
If `arrayName` is `pointer` , follow output : `&all element` 

* visit or evaluation for element
```go
arrayName[index]	// index is 0 ~ xxx
```

* visit (`refer to for or for-range` )
length of `array` : `len(arrayName)` 

### 3  memory
`各数组的元素在内存中连续存储` 
`差值` of `address` between `相邻` element   ==   size of `dataType` of element







## 2 `二维array` 
element of `array` is `array` 
concept is like `array` 
First `length of array` is which line
Second `length of array` is which column

### 3  `data type` format
```go
[length of arrary][length of element of array]dataType
```


### 3  define
```go
var arrayName [length of array][length of element of array]dataType	
```

### 3  evaluation
* not write `data type` of `variate` 
```go
arrayName = [length of array][length of element of array]dataType {{value, value ...}, {value, value ...} ...}
```

* not point `length of array` 
```go
arrayName = [...][length of element of array]dataType {{value, value ...}, {value, value ...} ...}
```
```go
arrayName = [length of element of array][...]dataType {{value, value ...}, {value, value ...} ...}
```
```go
arrayName = [...][...]dataType {{value, value ...}, {value, value ...} ...}
```

* point value of specific element by `index` 
```go
~~arrayName = [length of array][length of element of array]dataType {index:{index:value, index:value ...}, index:{index:value, index:value ...} ...}~~
```


### 3  define and    initialize or evaluation


### 3  use
* visit all element
```go
arrayName
```

* visit or evaluation for element
```go
arrayName[index]	// index is 0 ~ xxx
```

* visit or evaluation for element of element
```go
arrayName[index][index]	// index is 0 ~ xxx
```

* visit (`refer to for or for-range` )
length of `array` : `len(arrayName)` 
length of element of `array` : `len(arrayName[index])` 
`for-range` : `可以对value再进行for-range` 


### 3  memory
`各数组的元素在内存中连续存储` 
`各数组的元素的元素在内存中连续存储` 
`差值` of `address` between `相邻` element   ==   size of `dataType` of all element of element


### 3  `function` `参数传递` 
* `[length of array]arrayName` 

