##  define
`chan` , `<-chan` , `chan<-` : `都是同一种数据类型，他们之间的变量可以互相赋值，<-只是限制作用` 
we **must evaluation** before use variate

###   can read and write
```go
var chanVariateName   chan   dataType
```

###   only read
```go
var chanVariateName <-chan   dataType
```

###   only write
```go
var chanVariateName   chan<- dataType
```
