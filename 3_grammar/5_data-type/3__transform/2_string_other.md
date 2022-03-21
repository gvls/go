##  string    ->    other basic data type
If transform `"hello"` to `int`, result is default value of `int` 


###   package : strconv 
```go
boo, _ = strconv.ParseBool("true")
int64v, _ = strconv.ParseInt("111", 10, 64)
float64v, _ = strconv.ParseFloat(123.45, 64)
```


###   strconv.Atoi() 
short write of `strconv.ParseInt(s, 10, 0)` 
```go
strconv.Atoi()
```

