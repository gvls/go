## 2 `recursion` 
`递归的不断执行要不断接近递归出口` 
If not execute exit of `recursion` in the end, `memory` will `爆炸` 
It basic on `stack` 

```go
func <func name> (....) [...] {
	if ... {		// exit of recursion
		...
	}
	...
	[...] <func name>	// call itself
	...
}
```


