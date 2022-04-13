##  for range
It can visit `string` , `array` , `slice` , `channel` , `map` 
visit `map` is no order

* count of circulation for `array` or `slice` 
Before circulation, count of circulation already be calculate and it is fixation(length of array or slice)
when do add element in circulation, count of circulation is fixation

* count of circulation for `map` 
count of circulation map not is length of `map` 
when do add element for map in circulation, `key` and `value` **map** is new element

* `index` or `value` 
they is **new variate**(they is define by `:=` and define only do once) 
all `index` and `value` use **same** variate
value of they is **copy** from object

* If object is `string` 
type of `index` is `int` 
type of `value` is `rune`(`in32`)
So we will see some jump in `index` 

* If object is `[]XXX` 
type of `index` is `int` 
type of `value` is `XXX` 


###   format
```go
var str = "hello world"
for index, value := range str{	// name of index and value is custom. they is part variate
	...
}
```

```go
var ma = make(map[int]string)
for key := range ma{
	...
}
```

```go
var c = make(chan int, 10)
c <- 1
c <- 2
close(c)
for value := range c{
	...
}
```

