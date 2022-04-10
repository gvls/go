##  declaration initialize evaluation

###   declaration have type + 初始化 
can do `初始化` in `globe variate` 
```go
var <name> <type> = <value>
```

###   declaration not have type + 初始化
If not write `data type`, go will automatic conclude `data type` by value
It often be use in `globle` 
```go
var <name> = <value>
var <name1>, <name2> = 100, "tom"
var (
	<name1> = <value>
	<name2> = <value>
	<name3> <data type>
	...
)
```

###   do declaration and do evaluation
this way **Can't** write to `globle variate` (In globle, begin of statement must is identifier)
```go
<name> := <value>	// it   ==   var <name> dataType ; <name> = <value>
<name1>, <name2> := 100, "tom"
```
