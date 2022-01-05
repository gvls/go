## 2 `pointer`
A special **variate**
It save **address** of other variate or constant
Best to use `pointer` rather than `值类型` 
`pointer` can improve performance because needn't create create more space for `variate` and needn't use time to copy
When define `variate` of `pointer`, best make `variate name` is `pXxxx` 

## 2 Tips 
`变量加指针引用相当于套壳装箱: *<pointer> 相当于拆箱，转换为原来的变量` 

Think memory is `有很多存储格子的大柜子`
Door of every grid(`格子`) have a unique `门牌号`(`内存地址`)
`格子里面存的是值`
Think variate is grid
Make you focus on   value of grid (variate itself)   to track total process


## 2 distinguish`区分` 
```go
var ptr *<data type>
ptr		change track model (address)
*ptr		normal value model (variate and value)
&ptr		get address model (own address)
*(&ptr)
&(*ptr)
```
`<variate name>` get **own value**(address of other variate)
throw own grid content

`&<variate name>` get **own address**
throw own door number
`注意必要时候要在周围加括号为一个独立的运算` 

`*<variate name>` get **value** of memory of pointer point
Use value(door number) to find value of grid
It `约等于` other variate
`注意必要时候要在周围加括号为一个独立的运算，比如.会比*先运算` 
