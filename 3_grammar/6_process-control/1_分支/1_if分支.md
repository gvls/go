##  if 
It can be nest
divide one road to two or more road
procedure **only execute one** road of these road

* `<condition>` 
Can not write `()` 
Can include more than one statement and segmentation by `;` 
end statement : type of result is `bool` 
Can define variate in it

###   单分支 
best use this. Best use it to handle error
```go
if <condition> {
	...
}
```

###   双分支
not recommend
```go
if <condition> {
	...
} else {
	...
}
```

###   多分支 
not recommend
```go
if <condition> {
	...
} else if <condition> {
	...
}...
...
} else {
	...
}
```
