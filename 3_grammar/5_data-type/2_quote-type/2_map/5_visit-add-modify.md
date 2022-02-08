##  visit, add and modify
###   visit all key-value
output is **order** by `升序` in every time (after version 1.15.5 )
If `mapName` is `pointer` , follow output : `&all element` 
```go
mapName
```

###   visit, add, modify specific key-value
* common
```go
mapName[key]
```
If key not exist, it will return default value

* get value and judge if exist
```go
v, ok := mapName[key]
```

* value is `map` 
```go
mapName[key][key]
```

* visit all element
output is **not order** in every time
refer to `for-range` 
Because `key` may not is number and if is number, it may not `连续` so can't use `for` to visit


