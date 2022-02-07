##  value
`length of array` must same as length of `array variate`
`array` is **数值类型** , so `array` is use like `int` ...

###   default value
```go
[length]dataType{}
```

###   common
```go
[length]dataType{value, value ...}	// 读起来就像： 多少个 dataType 值分别是
```

###   not point **length** 
```go
[...]dataType{value, value ...}
```

###   point value of **specific element** by index 
```go
[length]dataType{
	index : value ,
	index : value ,
	...
}
```

###   get **pointer** 
add `&` to font of pre way

###   get **pointer** and evaluation in running
```go
new([length]dataType)
```

