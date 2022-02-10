##  MPG 
thread implement mode of Go
special two level

###   relation
```shell
				 KSE						KSE
kernel space	  ^
------------------|----------------------------
user space		  |
				  /\
				 /M \
				/____\
				  |
				|----|	/---\
				| P  |--| G	| \
				|____|	\---/  \
				  |		  |
				/---\	/---\
	running	--	| G	|	| G	| -- not running(runqueue)
				\---/	\---/
						  |
						/---\  /
						| G	| /
						\---/
```


###   M Machine 
one machine 关联了 one **kernel thread**
操作系统的主线程，物理线程 


###   P Processor 
count of it be set by GOMAXPROCS or runtime.GOMAXPROCS()
协程执行需要的上下文 
handler which handle user code
Processor == count of thread

* 全局 runqueue
When P finish all goroutine of itself, P can get goroutine from 全局 runqueue

* 均衡分配goroutine
When one P finish all goroutine of itself, it can get global goroutine or goroutine of other P


###   G goroutine 
light weight thread


