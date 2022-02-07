##  compiler in other system
When we in windows, we need add `set` in front of variate

|-------------|-------------------------|
| use way     | example                 |
|-------------|-------------------------|
| system path | `vim /etc/profile`      |
|-------------|-------------------------|
| one command | `GOOS=windows go build` |
|-------------|-------------------------|
such as :`GOOS` **can't use xiao xie**

###   Linux
```shell
CGO_ENABLED=0
GOOS=linux
GOARCH=amd64
go build
```

###   Windows
```shell
CGO_ENABLED=0
GOOS=windows
GOARCH=amd64
go build
```

###   Mac
```shell
CGO_ENABLED=0
GOOS=darwin
GOARCH=amd64
go build
```
