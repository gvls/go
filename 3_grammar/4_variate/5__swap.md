## 2 swap
If want to swap a and b

### 3  `并行赋值` 
```go
a, b = b, a

// it 等价:
// tmp := a
// a = b
// b = tmp
```

### 3  not use other `variate` 
```go
var a = 10
var b = 20
a = a + b
b = a - b
a = a - b
```
