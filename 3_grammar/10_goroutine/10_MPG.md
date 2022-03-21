##  MPG 
thread implement mode of Go
special two level

###   relation
```shell
				 KSE						KSE
kernel space	  ^
------------------|----------------------------
user space		  |
				  /\	/---\
				 /M \ --| G	| -- running
				/____\	\---/
				  |
				|----|	/---\
				| P  |--| G	| \
				|____|	\---/  \
				   		  |
				     	/---\
	                	| G	| -- not running(runqueue)
				     	\---/
						  |
						/---\  /
						| G	| /
						\---/

						/---\
						| G	| -- waiting
						\---/
```


###   M Machine 
one machine 关联了 one **kernel thread**
操作系统的主线程，物理线程 
bind one P


###   P Processor (context)
count of it be set by **GOMAXPROCS** or runtime.GOMAXPROCS()
defautl value of GOMAXPROCS is count of cpu
handler which handle user code
Max count is 256
goroutine need P when execute
Processor == count of thread
be save in global array
count of P decide coutn of M

* if have one free P
1. Go will **create one new M**
2. **bind one M** to it 
3. init one local runqueue to P



###   G goroutine 
**协程** not is 并发 but G support 并发
It is **user thread**, switch G under control of user
count of it be set by **GOMAXPROCS** or runtime.GOMAXPROCS()
light weight thread
When create one G , G will be add to local runqueue or global runqueue
**main()** also is one G

* search G fromt runqueue
local -> other local -> blobal

* local runqueue
belong to specific P

* global runqueue
When P finish all goroutine of itself, P can get goroutine from 全局 runqueue

* 均衡分配goroutine
When one P finish all goroutine of itself, it can get global goroutine or goroutine of other P
transmit count of local to local  : half
transmit count of local to global : count of global  /  count of P

####    switch G
When run procedure, Go will create **sysmon** to monitoring and manage
When G be execute once, sysmon will record the count
If one G be execute too long, sysmon will add tab to mark it
