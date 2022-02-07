##  visit field
###   visit all property
If `structName` is `pointer` , follow output : `&all element` 
```go
structName
```

###   visit field or evaluation
```go
structName.propertyName
```

* if variate is pointer
```go
(*variateName).property 
```
or
```go
variateName.property 		//go编译器会自动转换
```
