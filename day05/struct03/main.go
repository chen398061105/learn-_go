package main

import "fmt"

//结构体是值类型  type
type person struct {
	name string //字段名 + 字段类型 不能重复
	age  int
}

//构造函数：new名字 约定
//当结构体里面属性较多的时候，建议使用指针 减少内存开销
func newPerson(name string, age int) *person {
	return &person{
		name: name,
		age:  age,
	}
}

func main() {
	p := newPerson("小明", 11)
	fmt.Println(p)
}
