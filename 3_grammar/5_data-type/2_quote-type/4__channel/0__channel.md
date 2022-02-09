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
	buf					// save data. cache 循环link table
	qcount 
	dataqsiz
	elemsize
	closed
	elemtype
	sendx				// write index (channel <- data)
	sendq				// sudog queue 双向link table
	recvx				// 
	recvq				//
	lock		mute	// when do send or recv, this will lock current struct
}
```
```shell
buf
			read
			  ^
			  |
			  _________
			 /  |   |  \ 	\
		/	| 1	| 0 | 0 |
			|___________|	-
length		|   |   |   |
		\	| 2 |   | 0 |	-	capacity
			|___|___|___|
			|	|   |	|	-
	write ->| 0	| 0	| 0 |	
			 \_________/ 	/
       
1. add lock
2. **copy** data from goRountineMemory1/channel to channel/goRoutineMemory1
3. do unlock
```

