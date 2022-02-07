##  string    ->    other basic data type
If transform `"hello"` to `int`, result is default value of `int` 

###   package : strconv 
```go
boo, _ = strconv.ParseBool("true")
int64v, _ = strconv.ParseInt("111", 10, 64)
float64v, _ = strconv.ParseFloat(123.45, 64)
```
###   strconv.Atoi() 



##  other basic data type    ->    string
###   package : fmt
```go
str = fmt.Sprintf("%d", 111)
str = fmt.Sprintf("%f", 1.11)
str = fmt.Sprintf("%t", true)
str = fmt.Sprintf("%c", 'a')
```

###   package : strconv
```go
str = strconv.FormatInt(111, 10)
str = strconv.FormatFloat(1.111111, 'f', 4, 64)
str = strconv.FormatBool(true)
```

###   strconv.Itoa()
Data type of 111 must is `int` 
```go
str = strconv.Itoa(111)
```
