## 2 compiler in other system
When we in windows, we need add `set` in front of variate

### 3  Linux
```shell
CGO_ENABLED=0
GOOS=linux
GOARCH=amd64
go build
```

### 3  Windows
```shell
CGO_ENABLED=0
GOOS=windows
GOARCH=amd64
go build
```
example: `linux => windows: GOOS=windows go build` 

### 3  Mac
```shell
CGO_ENABLED=0
GOOS=darwin
GOARCH=amd64
go build
```
