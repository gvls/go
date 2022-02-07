## 2 `管道channel`
It is `queue` 
It is `线程安全的` : needn't add `lock` 
`data type` of all element of it is same



### 3  define
`chan` , `<-chan` , `chan<-` : `都是同一种数据类型，他们之间的变量可以互相赋值，<-只是限制作用` 

* can read and write
```go
var chanVariateName chan dataType
```

* only read
```go
var chanVariateName <-chan dataType
```

* only write
```go
var chanVariateName chan<- dataType
```



### 3  initialization
```go
chanVariateName = make(chan dataType, capacity)
```


### 3  use
must be initialization before use 

* get length of element
`len(chanVariateName)` 

* get capacity of `channel` 
`cap(chanVariateName)` 

* add element
When add element, if haven't free space to add, operator of add will `阻塞wait`    and if never have free space, `procedure` may appear `deadlock` 

```go
chanVariateName<- value
```

* remove element
When remove element, if haven't element to remove, operator of remove will `阻塞wait`    and if never have element, `procedure` may appear `deadlock` 
If `status` is `false` , `channel` is close and all data already read out

```go
variateName, status = <-chanVariateName
```

```go
<-chanVariateName
```

* close write
design `gorutime` must attention what time to close `channel` 
When close `channel` , we can read but can't write
When close `channel` and read out all data, read will `不再 阻塞wait` and return `false` 

```go
close(chaVariateName)
```

* exit `阻塞wait` 
All case will be calculate and random choose one `case` enter
In one `case`, If can't get value in `chaVariateName` , `procedure` not `阻塞wait` and execute next `case`
In one `case`, If can get value in `chaVariateName` , execute content of `case` 
If in all `case` can't get value in `chaVariateName` , execute content of `default` 
If in all `case` can't get value in `chaVariateName` and not have `default`, procedure will `阻塞wait` 
```go
select {
	case v := <-chaVariateName :
		...
	case v := <-chaVariateName :
		...
	...
	default :
		...
}
```


* visit
refer to `for` and `for-range` 
When remove element length must big than 0   otherwise it will `阻塞wait` and may appear `deadlock` 
`for-range` not have `index` in `return value`
If `channel` is close, `for-range` visit is normal
If `channel` isn't close, `for-range` will throw `deadock` error



