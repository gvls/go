##  实现了接口的数据类型 
not only is struct
Like `键盘` , `鼠标` 
one custom `daty type` can implement more than one `interface` 

* data type **have** method same as all method of interface
遵循了interface的标准 ==> interface的实现数据类型

###   format
```go
type intName int
type structName struct {
	...
}
func (typeVariateName [*]typeName) methodNameA(parameterList) returnList{
	...
}
```


* Attention is `dataType` or `*dataType` implement `method` of `interface` 
```go
type A interface{
	test()
}
type Stu struct {
}
func (s *Stru) test(){
	fmt.Println("test")
}
func main() {
	var s Stu 
	var a A = s		// error : `s` must be replace by `&s` 
	s.test()
}
```
