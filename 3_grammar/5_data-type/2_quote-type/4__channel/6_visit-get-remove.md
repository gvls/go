##  visit get remove
###   get and remove element
```go
variateName, status = <-chanVariateName
```
If `status` is `false` , `channel` is close and all data already read out


###   remove element
```go
<-chanVariateName
```


###   get and remove more
* for

* for-range
`for-range` for it not have `index` in `return value`
If can't get value from channel and channel not close, procedure will be **阻塞wait** 
If can't get value from channel and channel is close,  procedure will be **quit** for-range


###   monitoring more
* select
