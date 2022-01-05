## 2 use
Use must after define and evaluation (If value is `map` , notice value must evaluation before use)

* visit all
output is not order in every time (in version 1.15.5 : is `升序`)
If `mapName` is `pointer` , follow output : `&all element` 
```go
mapName
```

* visit value   or add value   or change to value
```go
mapName[key]
```

* visit value of value or evaluation to value of value or change to value of value
This type is value is `map` 
```go
mapName[key][key]
```

* delete `key-value` 
```go
delete(mapName, key)	// if key not exist, it not throw error
```

* delete all `key-value` 
visit all and delete
or `mapName = make ...` and old `map` will be `gc` recycle


* search `value` by `key` 
```go
value, findResult := mapName[key]	// type of findResult is bool
```

* visit all element
output is not order in every time
refer to `for-range` 
Because `key` may not is number, if is number, it may not `连续` , not use `for` to visit

 
* sort
get all `key` from `map` to `slice` , sort `slice` and visit `slice` to get `value` of `map` 

