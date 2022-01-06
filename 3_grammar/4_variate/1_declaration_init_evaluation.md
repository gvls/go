## 2 declaration initialize evaluation
value can is `function`

### 3  declaration
If declaration **part** `variate`(`struct variate` `除外`), we must use it   otherwise `build` will error
* `name` 
if `type` is `interface` : `*er` 
if have special word : such as : `APIClient` , `apiClient` , `userID` 

```go
var <name> <type>
var <name 1>, <name 2>, <name 3> <type>
```

* default initialize
If not evaluation `variate`, go will initialize it by default value
All `variate` will be initialize , include : `array` , `struct` ...
`int...` 0
`float...` 0.0
`complex...` 0
`string` ""
`bool` false
`quote type` nil (`还没有分配空间`)
`func...` nil
`error` nil



### 3  evaluation
```go
<name> = <value>
```



### 3  declaration + `初始化` 
can do `初始化` in `globe variate` 

#### 4   var + type
```go
var <name> <type> = <value>
```


#### 4   var
If not write `data type`, go will automatic conclude `data type` by value
It often be use in `globle` 
```go
var <name> = <value>
var <name 1>, <name 2> = 100, "tom"
var (
	<name 1> = <value>
	<name 2> = <value>
	<name 3> <data type>
	...
)
```



### 3  do declaration and do evaluation

#### 4   :
this way Can't write to `globle variate` 
```go
<name> := <value>	// it == var <name> dataType ; <name> = <value>
<name 1>, <name 2> := 100, "tom"
```
