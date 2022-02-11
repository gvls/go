##  select
**monitoring** and **visit** channel **once** 
All case will be calculate and **random** choose one case which execute success to enter
If channel of case is closed, this case is execute success

* process of total
If in all case execute failure, execute content of **default** 
If in all case execute failure, and not have default, procedure will **阻塞wait** 

* process of case
In one case, If execute failure, procedure not 阻塞wait and execute **next** case
In one case, If execute success, execute **content** of case 

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

* default
use it **careful** : not easy to judge what if channel is finish


###   performance
more channel will make select more slow
