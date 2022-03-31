##  context
It is a concurrence codeing mode 
It is a `interface` 
If superstratum is cancel, substratum will be cancel
If substratum is cancel, superstratum not be cancel
one task bind one context


```go
type Context interface{
	Deadline() (deadline time.Time, ok bool)
	Done() <-chan struct{}
	Err() error
	Value(key interface{}) interface{}
}
```

* Deadline()
传递结束信号以抢占并中断当前任务
return end time  which task be cancel 
If not have end time, ok == false

* Done()
指示一段时间后当前goroutine是否会被取消
If task is close, returen `closed channel`
It context is cancel, return `nil` 

* Err()
解释goroutine被取消的原因
If channel which Done() return isn't close, return `nil` otherwise return `not null value`
If context is close, return `Canceled`
If context is overtime, return `DeadlineExceeded`

* Value()
获取特定于当前任务树的额外信息
return value which key point to 
if key is not exist, returen `nil` 
