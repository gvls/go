##  waitGroup
A 并发safe struct

* goal
make main-goroutine wait other goroutine

###   define
```go
var wg sync.WaitGroup
```


###   add wait
```go
wg.Add(<number>)
```


###   complete one
```go
wg.Done()
```

```go
func xxx(..., done func()){
	defer done()
	...
}

func main() {
	...
	go xxx(..., wg.Done)
	...
}
```


###   wait complete
```go
wg.Wait()
```


