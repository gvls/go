##  close
If value of channel is nil, close will **panic** 
If channel is close, close will **panic** 
design `gorutime` must attention what time to **close** `channel` 
When close `channel` , we can read but **can't write**
```go
close(chaVariateName)	// only do once. if not may panic
```
When add element, if haven't free space to add, operator of add will **阻塞wait**    and if **never** have free space, procedure may appear **deadlock** 


###   tip
```shell
channel 就像一个仓库。write 就像生产者。 read 就像消费者。
write 不断添加商品进 channel。仓库满了就等待；若永远没人消费仓库一直满就死锁了；通知停产就不再等待添加；
read 不断从 channel 中运出商品。仓库空了就等待；若永远没新商品生产加入仓库据死锁了；通知停产就不再等待运出；
```


###   channel isn't close
####    write
* If have    free space
wrte new element to channel

* If haven't free space and may **have** goroutine do **read** for this channel
current goroutine will **阻塞wait** 

* If **haven't free** space and **never have** goroutine do **read** for this channel
procedure will **deadlock** 

####    read
* If have    element in channel
get and remove element from channel

* If haven't element in channel and may **have** goroutine do **write** for this channel
current goroutine will **阻塞wait** 

* If **haven't element** and **never have** goroutine do **write** for this channel
procedure will **deadlock** 


###   channel is close
####    write
**Can't** add any and **panic** 

####    read
* If have    element in channel
get and remove element from channel

* If haven't element in channel
**return** default value and false and do next code

