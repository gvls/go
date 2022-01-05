## 2 transform
When operate on different `data type`, we must do transform
`Go` can't transform data type **automatic**
Transform must `显式转换`
Object of transform is **value** of `variate` (`data type` of `variate` isn't changed)
```go
<data type>(<value>)
```

* `不用类型转换` 
`数字不管没有小数点 或者是'字符' 它可以是 int... 或者float...中的任何一个` 



## 2 type
### 3  Small   ->   big

### 3  Big   ->   small
#### 4   build out put error
By current `statement`, can know value is bigger than value scope of `data type` of `variate`

#### 4   Build not out put error
By current `statement`, can't know value is bigger than value scope of `data type` of `variate`
May have precision lost and value is value of overflow
Such as: 
```go
var v1 int64 = 130
var v2 int8 = int8(v1)
fmt.Println(v2)	//-126
```



## 2 string    ->    other basic data type
If transform `"hello"` to `int`, result is default value of `int` 

### 3  package : `strconv` 
```go
boo, _ = strconv.ParseBool("true")
int64v, _ = strconv.ParseInt("111", 10, 64)
float64v, _ = strconv.ParseFloat(123.45, 64)
```



## 2 other basic data type    ->    string
### 3  package : `fmt`
```go
str = fmt.Sprintf("%d", 111)
str = fmt.Sprintf("%f", 1.11)
str = fmt.Sprintf("%t", true)
str = fmt.Sprintf("%c", 'a')
```

### 3  package : `strconv`
```go
str = strconv.FormatInt(111, 10)
str = strconv.FormatFloat(1.111111, 'f', 4, 64)
str = strconv.FormatBool(true)
```

Data type of 111 must is `int` 
```go
str = Itoa(111)
```




## 2 `类型断言` 
```go
<variate>.(<type>)
```


