## 2 configuration
### 3  command
* get environment message
```go
go env [-json] [property]
```

* get help of `env` 
```go
go help env
```

* set `env` 
```go
go env -w <property>
```

* cancel `env` 
```go
go env -u <property>
```

#### 4   configure
```go
go env -w GOPROXY=https://goproxy.cn,direct
```


