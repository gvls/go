##  sudog
sudog include pointer which point to goroutine, data

```shell
type sudog struct{
	g 			*g
	isSelect	bool
	next		*sudog
	prev		*sudog
	elem		unsafe.Pointer

	acquiretime	int64
	releasetime	int64
	ticket		uint32
	parent		*sudog
	waitlink	*sudog
	waittail	*sudog
	c			*hchan
}
```


