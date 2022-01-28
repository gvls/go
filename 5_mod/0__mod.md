## 2 `mod` 
official manager for `package` 
Old `go` use `vendor` and `GOPATH` to manage `denpendency` 
When use `mod` , `dependency` also be download in `GOPATH/src/mod` and `go install` also save in `GOPATH/bin` 

### 3  file : `go.mod` 
annotation use `//` rather than `/* */` 

#### 4   add dependency
```go
require (
	...
)
```

#### 4   exclude dependency
```go
exclude (
	...
)
```

#### 4   replace dependency
```go
replace (
	...
)
```


