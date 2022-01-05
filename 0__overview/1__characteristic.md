### 3  function
`编码规范标准统一`
like: Go = c + python:
* have safe and performance of static language
* have develop maintain efficiency of dynamic language
one binary file : `可直接编译成机器码，不依赖其他库，glibc的版本有一定要求，部署就是`
`内置runtime`
`丰富的标准库，Go目前已经内置了大量的库，特别是网络库非常强大`
`完善的标准测试框架,单元，黑盒，白盒，压力测试等等`
`Go语言自带标准的性能分析工具，包括CPU、内存、阻塞操作(http请求，数据库请求，`
`跨平台编译，如果你写的Go代码不包含cgo，那么就可以做到window系统编译linux的应用`
`不一定需要IDE，任何地方写，任何地方运行`
high memory manage : `go内存占用更小，CPU消耗相对低`
`go基本上只需要标准库就足够了，实际工程中的性能表现要好于JAVA`
`Golang 默认一种垃圾回收策略，走的是高频次低延迟的路线，而默认Java/C#是高吞吐量高延迟的，所以这点非常不利于Go的性能表现。`
`面向接口编程` 

### 3  inherit c
Inherit many characteristic of c: have weak pointer
brief program language : less code can implement `框架的标准化` 

### 3  develop speed
`容易学习上手` 
have high develop speed

### 3  high compiler speed
have high compiler speed

### 3  package
Have package concept: one file is belong to one package and can't exist independent

### 3  garbage recovery
Automatic recovery rubbish
handle memory `泄漏` 

### 3  natural concurrence `并发`
Light weight concurrence
Handle more concurrence
Can use more core`核` efficiency
Basic on CSP concurrence model
`语言层面支持并发` 
use concurrence is easy 

### 3  `管道 channel`
implement communication between `goroutine`

### 3  more return valuate
Function have can have more return valuate

### 3  `切片 slice`

### 3  `延时执行 defer`
