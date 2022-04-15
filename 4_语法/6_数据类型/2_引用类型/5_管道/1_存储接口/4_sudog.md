##  sudog
保存 读取/写入协程的引用、需要写入的数据的地址/需要读入数据的空间的地址
```shell
type sudog struct{
	g 			*g
	isSelect	bool
	next		*sudog
	prev		*sudog
	elem		unsafe.Pointer

	acquiretime	int64
	releasetime	int64
	ticket		uint32
	parent		*sudog
	waitlink	*sudog
	waittail	*sudog
	c			*hchan
}
```
