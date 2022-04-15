##  csp
并发mode of Go : 多线程共享内存, csp
**Do not communicate by sharing memory; instead, share memory by communicating.**

###   common concurrence
like java
they communicating between 线程 by share memory
visit resource by lock
have 线程安全的 data struct
同一时间要限制读写变量的数量


###   csp concurrence
implement by goroutine and channel

####    goroutine 协程
unit of concurrence
like 线程

####    channel 管道
way of communication for goroutine
