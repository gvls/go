##  struct
```shell
logic :
type A struct{
	Name string
	Age int8
}

A{
	Name: "go",
	Age:  18,
}

physic :
   Name  Age
  /   \   |
_____________
|   |   |   |
|103|111| 18|
|   |   |   |
-------------
```
`结构内的属性顺序存储在内存中，但属性若是引用类型，它指向的地址不一定连续` 
It be use to create a `data type` 
It can include any `dataType` and number

