package main

import (
	"fmt"
)

func main() {
	//匿名结构体,多用于一次性使用
	var a = struct {
		name string
		age  int
	}{"xiaoming", 1}
	fmt.Println(a)
	p := newPerson("xiaomig", 18)
	p.say()
}

//方法和接收者 接收者：指的那个类型的变量可以调用这个函数
//　指针用法统一性，一个地方用了其他地方也要用
func (p *person) say() {
	fmt.Printf("%s :我是人的方法", p.name)
}

//构造函数
func newPerson(name string, age int) *person {
	return &person{
		name: name,
		age:  age,
	}
}

//结构体
type person struct {
	name    string
	age     int
	address //结构体嵌套
}
type address struct {
	province string
	city     string
}

//json序列化，
