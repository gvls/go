### 3  `TCP socket` 
`TCP socket` basic on `TCP/IP` 
Need `import "net"` 



#### 4   `server` 
1. monitor `port` 
2. receive `TCP` connection from `client` and build connection
3. create `goroutine` to handle `request` 

* create `Listener` 
```go
net.Listen(<protocol>, <socket>) (Listener, error)	// socket is 0.0.0.0:<port>
```

* close `listener` 
```go
<Listener>.close() error
```

* wait `connect` 
```go
for {
	<Listener>.Accept() (Conn, error)	// procedure will 阻塞 at here before receive connection
	...
}
```

* handle connection
```go
func <functionName>(conn net.Conn){
	defer conn.Close()			// close connection
	...
	for {
		buf := make([]byte, number)
		conn.Read(buf) (n int, err error)	// procedure will 阻塞 at here before receive data
		...
		str := string(buf[:n])
		...
	}
}
func main() {
	...
	go <functionName>(<Conn>)
	...
}
```



#### 4   `client` 
1. send connection to `server` 
2. send `data` to `server` and receive `data` from `server` 
3. close connection

* create `connection` 
```go
net.Dial(<protocol>, <socket>) (Conn, error)
```

* create input
```go
bufio.NewReader(os.Stdin) Reader
<Reader>.ReadString(<character of tail>) (string, error)
```

* send data
```go
<Conn>.Write([]byte) (int, error)	// int : number of byte
```

