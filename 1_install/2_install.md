## 2 install
### 3  Decompression
```shell
tar -zxvf xxxx.tar.gz
mv go/ /usr/local/go
```

### 3  environment
```shell
vim /etc/profile
	export GOROOT=/usr/local/go
	export PATH=$PATH:$GOROOT/bin
	export GOPATH=$HOME/go
	export PATH=$PATH:$GOPATH/bin
	export GO111MODULE=auto
	export GOPROXY=https://goproxy.cn,direct
	export GONOPROXY=git.tencent.com,git.code.tencent.com
go env -w GO111MODULE=on
// or export GO111MODULE=on (window use set GO111MOUDULE=on)
go env -w GOPROXY=https://goproxy.cn,direct
//go env -w GOPROXY=https://goproxy.io  or  https://proxy.golang.org
//go env -w GOSUMDB="off"

// cancel proxy
//go env -u GOPROXY
```

### 3  environment of `goland` 
```shell
settings -> Go ->
	GOROOT
	GOPATH
	GO Modules : GOPROXY=https://goproxy.io
```

### 3  Test
```shell
go version
```

