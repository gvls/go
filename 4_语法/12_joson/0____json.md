## 2 `json` 
It is a `轻量级` format of `数据交换` 
It is easy to read and write for people
It is easy be `解析` and `生成` by computer
It can improve efficiency of `网络传输` 
Any `data type` can be transform to `json` 

### 3  process
`language`  ==`序列化`==>  `json`  ===`网络传输`===>  `procedure`  ==`反序列化`==>  `other language`


### 3  format
`key-value` 
[check json if correct](https://www.json.cn/) 

```go
[
	{
		"key":value,
		"key":value,
		"key":[value,value],
		...
	},
	{
		...
	}
	...
]
```


### 3  `marshal序列化` 
Need `import "encoding/json"` 
`Property` of `struct` need `大写首字母` 

```go
json.Marshal(v interface{}) ([]byte, error)
// map, struct : {"key":value,"key":value}
// slice : [{"key":value},{"key":value}]
// int : value
```


### 3  `unmarshal反序列化` 
Need `import "encoding/json"` 
`variate` of `map`, `slice` can not evaluation or initialize (`底层自动make`)

```go
var s dataType		// use json and &s to 调用 follow function
json.Unmarshal([]byte, *dataType) error
// dataType of varaite need same as dataType which before marshal
```
