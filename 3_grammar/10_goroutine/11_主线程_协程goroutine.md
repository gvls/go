##  `主线程` | `进程` 
Like `进程` , but `比进程轻量` 
It is physics `线程` and is `重量级` 

##  `协程` | `线程` 
Like `线程` , but `比线程轻量` 
`一个线程可以有多个协程` 
It is start by `主线程`
It be create as `stack` in `heap` `来模拟线程` 
It is logical `线程` and is `轻量级` 
`主线程退出时所有协程也会退出` 
`Go` can easy to open `上万` `goroutine` 
goroutine : user space light weight thread
use less memory than thread
have high performance on call

协程（coroutine）：又称微线程与子例程（或者称为函数）一样，协程（coroutine）也是一种程序组件。相对子例程而言，协程更为一般和灵活，但在实践中使用没有子例程那样广泛。

和线程类似，共享堆，不共享栈，协程的切换一般由程序员在代码中显式控制。它避免了上下文切换的额外耗费，兼顾了多线程的优点，简化了高并发程序的复杂。

Goroutine和其他语言的协程（coroutine）在使用方式上类似，但从字面意义上来看不同（一个是Goroutine，一个是coroutine），再就是协程是一种协作任务控制机制，在最简单的意义上，协程不是并发的，而Goroutine支持并发的。因此Goroutine可以理解为一种Go语言的协程。同时它可以运行在一个或多个线程上。
1、开销小

POSIX的thread API虽然能够提供丰富的API，例如配置自己的CPU亲和性，申请资源等等，线程在得到了很多与进程相同的控制权的同时，开销也非常的大，在Goroutine中则不需这些额外的开销，所以一个Golang的程序中可以支持10w级别的Goroutine。

每个 goroutine (协程) 默认占用内存远比 Java 、C 的线程少（*goroutine：*2KB ，线程：8MB）

2、调度性能好

在Golang的程序中，操作系统级别的线程调度，通常不会做出合适的调度决策。例如在GC时，内存必须要达到一个一致的状态。在Goroutine机制里，Golang可以控制Goroutine的调度，从而在一个合适的时间进行GC。

在应用层模拟的线程，它避免了上下文切换的额外耗费，兼顾了多线程的优点。简化了高并发程序的复杂度。

缺点：

协程调度机制无法实现公平调度。




###   create one `goroutine` 
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
