## 2 character
Use `int..` or `float..` to save
Actually save number which char mapper (`Go` use `UTF-8`)
So we can calculate the value by `operator`

## 2 value
We can get value from UTF-8 table or ASCII table (`字符编码表`)
```go
'0' == 48
'A' == 65
'a' == 97
```

## 2 `\`
```go
'\n'	// is switch next line char
'\a'	// loud
'\b'	// back `格`
'\f'	// switch page
'\r'	// back car   and string after `\r` will override appear in begin of line
'\t'	// make table and align
'\v'	// make vertical table
'\''	// `'` 
'\"'	// `"` 
'\\'	// `\` 
```


## 2 process
store:
character -> code value -> binary -> store

Read:
binary -> code value -> character -> read
