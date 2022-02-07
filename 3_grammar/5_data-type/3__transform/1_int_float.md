##  int and float
###   Small   ->   big

###   Big   ->   small
* build out put error
By current `statement`, can know value is bigger than value scope of `data type` of `variate`.

* Build not out put error
By current `statement`, can't know value is bigger than value scope of `data type` of `variate`.
May have **precision lost** and value is value of **overflow**
Such as: 
```go
var v1 int64 = 130
var v2 int8 = int8(v1)
fmt.Println(v2)	//-126
```
