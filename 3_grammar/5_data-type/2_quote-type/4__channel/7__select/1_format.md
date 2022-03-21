##  format
###   monitoring specific
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
* `case`
**must** do send or receive on channel
If channel of case is **closed**, this case is execute success

* `break`
will jump out of `select` rather than `for` 

* `default` 
use it **careful** : not easy to judge what if channel is finish



###   monitoring all
always wait
```go
select{}
```
