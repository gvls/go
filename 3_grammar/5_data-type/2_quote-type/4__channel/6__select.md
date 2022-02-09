##  select
* exit `阻塞wait` 
All case will be calculate and random choose one `case` enter

If in all `case` can't get value in `chaVariateName` , execute content of `default` 
If in all `case` can't get value in `chaVariateName` and not have `default`, procedure will `阻塞wait` 

In one `case`, If can't get value in `chaVariateName` , `procedure` not `阻塞wait` and execute next `case`
In one `case`, If can get value in `chaVariateName` , execute content of `case` 

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

