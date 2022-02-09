##  swap
If want to swap a and b

###   并行赋值 
```go
a, b = b, a

// it 等价:
// tmp := a
// a = b
// b = tmp
```

###   not use other variate 
```go
var a = 10
var b = 20
a = a + b
b = a - b
a = a - b
```
