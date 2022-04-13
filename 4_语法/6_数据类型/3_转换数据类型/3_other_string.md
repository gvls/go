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
short write of `strconv.FormatInt(i, 10)` 
```go
str = strconv.Itoa(111)
```
