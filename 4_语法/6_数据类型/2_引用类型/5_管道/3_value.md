##  value
we **must evaluation** before use variate
If value is **nil** , visit and add will waiting
If value is **nil** and use in **select** , select never visit data from this channel

```go
make(chan dataType, capacity)
```
