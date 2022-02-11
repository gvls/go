##  memory
* 先进先出
先从 Channel 读取数据的 Goroutine 会先接收到数据
先向 Channel 发送数据的 Goroutine 会得到先发送数据的权利

```shell
type hchan struct {
	buf			unsafe.Pointer	// buffer cache data.
	qcount 						// count of circulation
	dataqsiz					// length of circulation
	elemsize					// enable max size of element
	closed						// flag : channel if closed
	elemtype					// type of element

	sendx						// write index(write pointer) of buf (channel <- data)
	f 
	recvx						// 

	sendq		waitq			// waiting queue
	recvq		waitq			//
	lock		mutex			// 互斥锁 when do send or recv for buf, this will lock current struct
}
```
