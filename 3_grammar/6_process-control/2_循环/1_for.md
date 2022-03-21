##  for 
It can visit `string` , `array` , `slice` 

* `expression 1`
only execute once and execute **before first enter** `for` 

* `condition`
execute **before per enter** `for` 
`condition` is `expression` which return `bool` 

* `expression 2`
execute **after per exit** `for` 


###   format 
Visit `variate` by size of element of `data type` 

```go
for [expression 1] ; [condition] ; [expression 2] {
	...
}
```

```go
[expression 1]
for <condition> {	
	...				// write break in first line can implemnt while
	...
	...
	[expression 2]	// write break in end line can implement do-while
}
```

```go
[expression 1]
for {				// for {    ==    for ; ; {
	...				// write break in first line can implemnt while
	...
	...
	[expression 2]	// write break in end line can implement do-while
}
```
