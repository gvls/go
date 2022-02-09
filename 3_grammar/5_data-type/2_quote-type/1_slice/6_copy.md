##  copy
###   浅拷贝
copy **address** 
array of slice1 **is** array of slice2
```go
slice1 = slice2 		// slice1  <==value==  slice2
```
```go
func funcName(s []datType){	// s is new variate and value of pointer is old array
	...
}
func main(){
	...
	funcName(s)
}
```

###   深拷贝
copy **value of array** which pointer point to
length and capacity **isn't change**
array of slice1 is **new** array
```go
copy(slice1, slice2)	// slice1  <==value==  slice2
```

* if len1  <  len2 
`array1` may **lose** some element of `array2`
So slice1 **must** be initialize

* if len1  >  len2 
end element of `array1` keep old value


