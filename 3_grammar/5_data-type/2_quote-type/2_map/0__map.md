## 2 map
It is a `key-value` data structure
It's `存储` is `无序` 
Capacity of `map` is dynamic increase
Element of `map` can't be `取地址` 


### 3  `key` 
can't repetition
If `key` which want to add same as `key` already in `map` , new `key-value` will overwrite old `key-value`

### 3  `data type` 
* `keyType` 
Can't is `slice` , `map` , `function` (they can use `==` to judge
Usually, `keyType` is `int` or `string` 

* `valueType` 
Can is `map` , `struct`(often use) ...

```go
map[keyType]valueType	//读起来就是 map[key] 的值是 value
```

### 3  define
```go
var variateName map[keyType]valueType
```

### 3  evaluation
```go
variateName = map[keyType]valueType {
	key : value,
	...
	key : value,
}
```

```go
variateName = make(map[keyType]valueType [, size]) // if not write size, size == 1
```

