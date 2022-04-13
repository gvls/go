## 2 command
Need be do in `$GOPATH` 

### 3  build project
`go build` 

### 3  download `module` to local cache
path is `$GOPATH/pkg/mod/cache` 
```go
go mod download
```

### 3  edit `go.mod` 
```go
go mod edit ...
```

### 3  display picture of dependency of `module` 
```go
go mod graph
```

### 3  initialize
create `module` (`go.mod`)
```go
go mod init [path/name]
```

### 3  add missing `package` 
```go
go mod tidy
```

### 3  copy dependency to `vender/` 
```go
go mod vender
```

### 3  confirm dependency
check dependency and view if change
```go
go mod verify
```

### 3  explain why need `package` 
```go
go mod why
```

### 3  see all dependency of `module` 
```go
go list -m all
```

### 3  update stabled dependency
```go
go get rsc.io/samler
```

### 3  update specific version of dependency
```go
go list -m -version rsc.io/sampler
go ger rsc.io/sampler@v1.3.1
```

### 3  ignore `package` in `cache` and only use dependency in `vendor` 
```go
go build -mod=vendor
```

### 3  download dependency to `GOPATH/pkg/mod` and write dependency to `go.mod` 
```go
go get -m <package>
```

### 3  update all dependency of main `package` 
```go
go get -u=patch
```


