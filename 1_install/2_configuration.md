##  install
###   Decompression
```shell
tar -zxvf <xxxx>.tar.gz
mv go/ /usr/local/go
```

###   environment
1. profile 
```shell
vim /etc/profile
	export GOROOT=/usr/local/go
	export PATH=$PATH:$GOROOT/bin
	export GOPATH=$HOME/go
	export PATH=$PATH:$GOPATH/bin
	export GO111MODULE=auto
	export GOPROXY=https://goproxy.cn,direct
	export GONOPROXY=git.tencent.com,git.code.tencent.com
```
2. go mod 
```shell
# get environment message
#go env [-json] [property]

# set `env` 
#go env -w <property>

# cancel `env` 
#go env -u <property>

go env -w GO111MODULE=auto
# or export GO111MODULE=on
# (window use set GO111MOUDULE=on)

go env -w GOPROXY=https://goproxy.cn,direct
#go env -w GOPROXY=https://goproxy.io  or  https://proxy.golang.org

#go env -w GOSUMDB="off"

# cancel proxy
#go env -u GOPROXY
```

###   environment of goland 
```
settings -> Go ->
	GOROOT
	GOPATH
	GO Modules : GOPROXY=https://goproxy.io
```

###   Test
```shell
go version
```

