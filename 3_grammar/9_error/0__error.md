## 2 error 
When part of procedure throw error, procedure will **exist**

### 3  `error` 
It is a `data type`
It can catch error and do somethings 
procedure continue running

`error` == `nil` : result is success
`error` != `nil` : result is false

#### 4   `defer` 

#### 4   `panic` 

#### 4   `recover` 


### 3  custom `error` 
need `import "errors"` 

```go
errors.New("<cusom error>")	// return error
// use `panic` to output error and exist procedure
```

