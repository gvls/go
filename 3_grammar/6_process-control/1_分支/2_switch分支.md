## 2 `swirch` 
divide one road to two or more road
procedure **only execute one** road of these road
like `if-else-if-else` 
From top to bottom, search which `expresstion` == `expresstion 0` 

* `expresstion` 
it is `constant` , `variate` , `function` which have return ...
Type of all `expresstion` must same
in `constant` , `expresstion` can't same

* `fallthrough` 

* according to value
```go
switch <expresstion 0> {
	case <expresstion 1>, <expresstion 2> [, ...] :
		...
	case <expresstion 3>, <expresstion 4> [, ...] :
		...
	[default :]
		...
}
```

* according to `bool` 
If not write `expression 0`   or define `variate` by write `;` to end(not recommend), can use `condition` to replace `expresstion`  and use `switch` as `if-else-if-else` 
```go
switch {
	case <condition 1>[, ...] :
		...
	case <condition 3>, <condition 4> [, ...] :
		...
	[default :]
		...
}
```

* `Type switch` 
judge actual type of `variate` of `interface` 
```go
switch 	interfaceVaraite.(type) {
	case <type 1> [, ...] :
		...
	case <type 3>, <type 4> [, ...] :
		...
	[default :]
		...
}
```

