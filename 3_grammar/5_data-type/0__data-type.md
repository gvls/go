## 2 `值类型` 
Use one block memory
Memory store value directly
often be create in stack
After define `variate`, we can visit or evaluation for element of it
`变量和变量之间 可以 用==直接判断` 
```go
int   float   bool   string   array   struct
```



## 2 `quote type` 
Use two block memory
Memory store address of actual value
Often be create in heap, when if don't quote any actual value, `GC` will recycle this garbage
When define `variate` , we must visit element of it   before evaluation for `variate` 
`变量和变量之间 不可以 用==直接判断` 
```go
pointer   slice   map   interface   管道channel
```



## 2 `custom type` 
Create a **new** `data type` according to specific `data type` 
Need handle `scope of effect` : may define in `globe` 
```go
type <name of custom type> <data type>
```

* follow is error :
```go
type myInt int
var a myInt
var b int
a = b
```



## 2 `function` 
