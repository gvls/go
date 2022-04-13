##  buf
data struct is circulation link table (先进先出)

* size of buf
in `make(chan dataType)`, if size not write,  size == 0 by default
If size == 0, this channel not have buf

* full
one space isn't used
recvx == **sendx + 1** (recvx is head, sendx is tail)

* empty
sendx == recvx

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
       
```


###   process
1. add lock
2. **copy** data from goRountineMemory1/channel to channel/goRoutineMemory1
3. do unlock


###   不完整的 not lock queue
* If have buf
It is like **async异步**
If not full and not empty may use buf
If full or empty may use sendq and recvq

* If not have buf
It is like **sync同步**
It **like** what channel have buf **but** full or empty
send and recv only use sendq and recvq
send and write need **waiting**

* chan struct{}
It is like **async异步**
`struct{}` not use space of memory
neen't implement buf and send direct

