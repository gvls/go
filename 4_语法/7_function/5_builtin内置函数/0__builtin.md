## 2 `builtin`
When use it, we needn't `import package`

### 3  get length
`len()` 
Get length by **byte** : `string` 
Get length by **element** : `arrary` , `slice` , `channel通道`(`未读取`) , `map映射` , `数组指针` 
Return value is `int` 


### 3  allocate memory for `值类型` 
`new(<data type>)   * data type` 
Return value is `pointer` 
It `有点像` `var      * data type` 

### 3  allocate memory for `引用类型` 
`make(<data type> ... )   data type` 
It `有点像` `var      data type` 

### 3  get capacity
`cap()` 

### 3  out put binary
`unsafe.Sizeof(<name>)` 
out put the binary of name
