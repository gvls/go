##  value
###   default value
```go
structName{}
```

###   common
must write value of **all field** 
```go
structName{value1, value2 ...}
```

###   **point field** for value
can write value of part field 
```go
structName{
	property1 : value1 ,
	property2 : value2 ,
	...
}
```

###   get **pointer** 
add `&` to font of pre way

###   get **pointer** and evaluation in running
```go
new(structName)
```
