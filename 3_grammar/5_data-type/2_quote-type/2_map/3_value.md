##  value
we **must evaluation** before use variate

* key
can't repetition
If new `Key` which want to add same as `key` already in `map`, new `key-value` will overwrite old `key-value`

###   空map
```go
map[keyType]valueType{}
```

###   common
```go
map[keyType]valueType {
	key : value,
	...
	key : value,
}
```

###   pointer
add `&` to front of pre way

###   in running to new array 
```go
make(map[keyType]valueType [, size]) // if not write size, size == ~~ 1 ~~
```

