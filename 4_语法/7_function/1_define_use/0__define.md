## 2 define
if more than one `parameter type` of `parameter` is same, can only write `parameter type` after last `parameter` 

* point `return value` in use `return` 
```go
// if only have one return value, can not write `()`  
func <name>([<parameter name> <parameter type>[, ...]])   [(<return type>[, ...])]{
	...
	[return [return value][, ...]]	
	...
}
```

* automatic point `return value` in use `return` 
```go
// `variate` of `return name` is already be create
// `return list` must add `()` to two side
func <name>([<parameter name> <parameter type>[, ...]])   [(<return name> <return type>[, ...])]{
	...
	[return]	
	...
}
```
