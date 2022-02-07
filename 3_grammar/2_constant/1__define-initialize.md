##  define initialize
`dataType` : `bool` , `int...` , `float...` , `complex` , `string` 
`value` can't include `variate` or `function` ...

* define one in one line
```go
const <name> [dataType] = <value>	// name all big case and connected by _
```

* define more in one line
```go
const <name>, <name>, <name> = <value>, <value>, <value> // name ditto
```

* define more in more line
```go
const (
	<name> = <value>
	<name> = <value>
	...
)
```

* `iota` 
```go
const (
	<name1> = iota		// first line iota is 0
	<name2>				// value is value of pre lien. per pass one line iota++ 
	<name3>				// ditto
	...
)
```
