## 2 `string`
```shell
"5201314"
_____________________________
|   |   |   |   |   |   |   |
| 53| 50| 48| 49| 51| 49| 52|
|   |   |   |   |   |   |   |
-----------------------------
```
Go use UTF-8 to `标识` Unicode text
`底层` is `[...]byte` : it can be `slice` and can't visit china character by element
If have `"` in `""` , `"` need write `\"` 

* `string` is constant 
it can't be changed again
Such as: 
```go
var str = "hello"
str[0] = 'H'	// error
```

* get length of `string` 
```go
len(<string>)
```



## 2 `表现形式` 
* `双引号 ""`
```go
"<content>"
```

* `反引号 `
output string by `原生` way
output can include: special character, source code  (prevent attack)
```go
`<content>` 
```



## 2 `拼接`joint
Use symbol : `+` 

Example:
```go
var str = "hello" +
	"world"
```

## 2 visit
* visit element by byte (can't visit china)
```go
`<string>[<index>]` // `index` is begin from `0` 
```

* visit element by character (can visit china)
`refer to for-range` 
`transform to []rune` 


## 2 change value of `string` (not is change variate)
1. `string` => `[]byte`(`处理字母和数字`)  or  `string` => `[]rune`(`处理中文`) 
`arrayName` := `[]byte(stringName)`  or  `arrayName` := `[]rune(stringName)` 
2. change element
3. `[]byte` => `string`  or  `[]rune` => `string`  
`stringName` = `string(arrayName)`  or  `stringName` = `string(arrayName)` 
