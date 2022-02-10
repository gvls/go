##  管道channel
It is queue (先进先出)
It is **线程安全的** : needn't add lock 
`data type` of all element of it is same

* get length of element
`len(chanVariateName)` 

* get capacity of channel 
`cap(chanVariateName)` 


```shell
type hchan struct {
	buf			unsafe.Pointer	// save data. cache circulation link table
	qcount 						// count of circulation
	dataqsiz					// length of circulation
	elemsize					// enable max size of element
	closed						// flag : channel if closed
	elemtype					// type of element

	sendx						// write index of buf (channel <- data)
	f 
	recvx						// 

	sendq		waitq			// sudog struct queue 双向link table
	recvq		waitq			//
	lock		mutex			// when do send or recv, this will lock current struct
}
// sendq be use in what buf is full
// recvq be use in what buf is empty
```
```shell
buf
				recvx
				  ^
				  |
				read
				  |
				  _________
				 /  |   |  \ 	\
			/	/ 1	| 0 | 0 \
				|___/___\___|	-
	length		|   |   |   |
			\	| 2 |   | 0 |	-	capacity
				|___\___/___|
				|	|   |	|	-
sendx---write-->\ 0	| 0	| 0 /	
				 \_________/ 	/
       
1. add lock
2. **copy** data from goRountineMemory1/channel to channel/goRoutineMemory1
3. do unlock
```


###   send阻塞
1. GoroutineA do send(lock)
2. check if full for buf
3. if full : **goroutineA** 让出 M and 主动 call Go caller to **sleep**
4. **sudog** struct of GoroutineA will be create and will be save in **sendq** and **sleep**
5. other goroutine get M and execute
6. ...
7. goroutineB do recv(lock) from **buf** and buf will be not full
8. **sudog** of **sendq** will be **active** and send to buf
9. **active** **goroutineA** and make it to goqueue to wait be execute


###   recv阻塞
1. GoroutineB do recv(lock)
2. check if empty for buf
3. if empty : **goroutineB** 让出 M and 主动 call Go caller to **sleep**
4. **sudog** struct of GoroutineB will be create and will be save in **recvq** and **sleep**
5. other goroutine get M and execute
6. ...
7. goroutineA do send to **goroutineA** 
8. **sudog** of **sendq** will be **active** and recv from cache
9. **active** **goroutineA** and make it to goqueue to wait be execute
