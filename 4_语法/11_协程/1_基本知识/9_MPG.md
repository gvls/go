##  MPG 
Go 的线程实现模型

###   关系
```shell
				 KSE						KSE
内核空间		  ^
------------------|----------------------------
用户空间		  |
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
一个`M` 关联 一个内核线程
操作系统的主线程，物理线程 
bind one P


###   P Processor (context)
最大数量是 256
`P` 决定`M` 的数量，`P` 的数量 == `M` 的数量
在全局数组中保存
`G` 的执行需要`P` 
handler which handle user code

* 若有一个空闲的`P` 
1. Go 会创建一个新的`M` 
2. 绑定`M` 给`P`  
3. 初始化一个本地`runqueue` 给`P`  



###   G goroutine 
**协程** 不是 并发 但 G 支持 并发
轻量级的用户线程，`G` 的切换由用户控制
当`G` 被创建后，会被添加到本地`runqueue`或全局`runqueue` 
`main()` 也是一个`G` 

* 本地`runqueue` 
属于特定的`P` 

* 全局`runqueue` 
供所有`P`使用

* 均衡分配`G` 
当`P` 完成了自己本地的`runqueue` ，它可以从其他`P` 的本地`runqueue` 或者全局`runqueue` 获取`G` 
从其他`P` 的本地`runqueue` 获取的数量：一半
从全局`runqueue` 获取的数量：全局`runqueue` 的`G` 数量 / `P` 的数量

####    switch G
When run procedure, Go will create **sysmon** to monitoring and manage
When G be execute once, sysmon will record the count
If one G be execute too long, sysmon will add tab to mark it
