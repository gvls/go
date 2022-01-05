## 2 define initialize
`dataType` : `bool` , `int...` , `float...` , `complex` , `string` 
`value` can't include `variate` or `function` ...

### 3  define one
```go
const <name> [dataType] = <value>
```

```go
const <name>, <name>, <name> = <value>, <value>, <value>
```


### 3  define more 
```go
const (
	<name> = <value>
	<name> = <value>
	...
)
```

### 3  `iota` 
`第一次出现是0, 每经过一行statement 他++` 

### 3  define and not write value
`若当前行不写值，值是上一行的值` 

```go
const (
	<name1> = iota		// value(这一行的值) is 0
	<name2>			// value(这一行的值) is name1(第一行的值) + 1
	<name3>			// value(这一行的值) is name1(第一行的值) + 1 + 1
	...
)
```
