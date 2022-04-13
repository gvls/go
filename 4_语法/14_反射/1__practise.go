//## 2 `练习`
//`通过反射遍历结构体的字段，调用结构体的方法，并获取结构体标签的值`
package main

import (
	"fmt"
	"reflect"
)

type Monster struct {
	Name  string `json:"name"`
	Age   int    `json:"monster_age"`
	Score float32
	Sex   string
}

func (s Monster) Print() {
	fmt.Println("start")
	fmt.Println(s)
	fmt.Println("end")
}
func (s Monster) GetSum(n1 int, n2 int) int {
	return n1 + n2
}
func (s Monster) Set(name string, age int, score float32, sex string) {
	s.Name = name
	s.Age = age
	s.Score = score
	s.Sex = sex
}

func show(a interface{}) {
	var typ = reflect.TypeOf(a)
	var val = reflect.ValueOf(a)
	var kin = val.Kind()
	if kin != reflect.Struct {
		fmt.Println("expect struct")
		return
	}
	num := val.NumField()
	fmt.Printf("struct has %d fields\n", num)
	for i := 0; i < num; i++ {
		fmt.Printf("Field %d:\tname\t%v\t", i, val.Field(i))
		tagVal := typ.Field(i).Tag.Get("json")
		if tagVal != "" {
			fmt.Printf("%v", tagVal)
		}
		fmt.Println()
	}
	numOfMethod := val.NumMethod()
	fmt.Printf("\nstruct has %d methods\n", numOfMethod)

	fmt.Printf("\nmethod 1:\n")
	var params []reflect.Value
	params = append(params, reflect.ValueOf(10))
	params = append(params, reflect.ValueOf(40))
	res := val.Method(0).Call(params)
	fmt.Printf("res= %d\n", res[0].Int())

	fmt.Printf("\nmethod 2:\n")
	val.Method(1).Call(nil)
}

func main() {
	var a Monster = Monster{
		Name:  "fox",
		Age:   300,
		Score: 30.8,
	}
	show(a)
}
