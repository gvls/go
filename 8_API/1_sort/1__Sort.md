## 2 sort
* sort `interface` 

```go
type Interface interface {
	Len() int		// length
	Less(i, j int) bool	// compare method. When result is true, 按i<j 还是i>j
	Swap(i, j int)		// sway i and j
}

sort.Sort(<Interface>)
```


