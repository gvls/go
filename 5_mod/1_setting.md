## 2 setting
```shell
go env -w ...
```

### 3  `GO111MODULE` 
After `go` version 1.11, `GO111MODULE=auto` 
* `on` 
use `go module` and ignore `GOPATH` and `vendor` 

* `off` 
search `vendor` or `GOPATH` 

* `auto` 
If current directory isn't in `$GOPATH` **and** current directory or father directory have `go.mod` , use `GO111MODULE` and ignore `GOPATH`   otherwise use `GOPATH` 
