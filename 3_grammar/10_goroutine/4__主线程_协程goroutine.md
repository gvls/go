## 2 `主线程` | `进程` 
Like `进程` , but `比进程轻量` 
It is physics `线程` and is `重量级` 

## 2 `协程` | `线程` 
Like `线程` , but `比线程轻量` 
`一个线程可以有多个协程` 
It is start by `主线程`
It be create as `stack` in `heap` `来模拟线程` 
It is logical `线程` and is `轻量级` 
`主线程退出时所有协程也会退出` 
`Go` can easy to open `上万` `goroutine` 




### 3  create one `goroutine` 
If one `goroutine` is throw `panic`, total `procedure` will exit (can use `defer` and `recover` to handle)
```go
go <function>
```




### 3  `资源竞争` 
We can add `-race` to `go build` to check `是否有资源竞争` 

#### 4   add lock
Like `上厕所关门，门关只能在外面排队列等待` 
Need `import "sync"` 
suit for small `concurrence` 
But `主线程不知道要等多久，协程` 

* define `互斥锁` in `globle` 
```go
var variateName sync.Mutex
```

* use
```go
variateName.Lock()
// operator of write or read on 共同竞争的资源
variateName.Unlock()
```

#### 4   `channel` 
suit for big `concurrence` 
1 ~ more `channel` be use to `共享数据` 
define `flag`(this may is `竞争资源`) to judge if exist (`主线程一直for判断这个flag`)   or define `channel` to receive result of all `goroutine` and `主线程` remove number of `goroutine` element from `channel` 




### 3  feature
`独立的栈空间` 
`共享程序堆空间` 
`调度由用户控制` 
`轻量级的线程` 
