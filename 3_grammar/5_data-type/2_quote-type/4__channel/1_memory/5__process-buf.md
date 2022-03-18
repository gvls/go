##  process by buf
###   send
1. do send
2. If channel is **close** , return result 
3. If exist **waiter**, send data to waiter direct (buf is empty)
4. If buf have free space, send data to buf 
5. or else : waiting what other goroutine do recv 

* send direct
1. copy data to first waiter
2. make waiter is preparing (move to runqueue)
3. waiter preparing **wait** be execute

* send to buf
1. do send
2. calculate index of free space of buf
3. send data to buf and update sendx and qcount 

* waiting to send
1. get goroutine and data of sender
2. create sudog 
3. add sudog to sendx 
4. set sender is waiting (wait for sudog is active)
5. if use **select** sender not waiting 


###   recv
1. do recv
2. If channel is **close** and buf is empty, return result
3. If exist **waiter**, recv data from waiter direct 
4. If buf have data, recv data from buf 
5. or else : waiting what other goroutine do send 

* recv direct
1. If not have buf, copy from first waiter
2. If have buf,		copy data from buf and copy data of first waiter to buf
3. set waiter is preparing

* recv from buf
1. calculate index of first element of buf
2. get data from buf and update sendx and qcount 

* waiting to recv
1. get goroutine and data of recver
2. create sudog 
3. add sudog to recvq 
4. set recver is waiting (wait for sudog is active)
5. if use **select** recver not waiting 

###   send阻塞
in total process, M(system thread) isn't **waiting** (use a little resource switch goroutine)

1. GoroutineA do send(lock channel)
2. check if full for buf
3. if full : **goroutineA** 让出 M and 主动 call Go caller to **waiting**
4. **sudog** struct of GoroutineA will be create and will be save in **sendq** and **waiting**
5. goroutineB of runqueue get M and execute
6. ...
7. goroutineB do recv(lock channel) from **buf** and buf will be not full
8. **sudog** of **sendq** will be **active** and send to buf
9. **active** **goroutineA** and make it to runqueue to wait be execute


###   recv阻塞
in total process, M(system thread) isn't **waiting** (use a little resource switch goroutine)

1. GoroutineB do recv(lock channel)
2. check if empty for buf
3. if empty : **goroutineB** 让出 M and 主动 call Go caller to **waiting**
4. **sudog** struct of GoroutineB will be create and will be save in **recvq** and **waiting**
5. goroutineA of runqueue get M and execute
6. ...
7. goroutineA do send to **goroutineB** 
8. **sudog** of **recvq** will be **active** and recv from cache
9. **active** **goroutineB** and make it to runqueue to wait be execute
