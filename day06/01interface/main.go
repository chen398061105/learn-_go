package main

import "fmt"

//接口是一种类型，约束变量的方法,该变量的方法必须实现接口的所有方法
func main() {
	var d1 dog
	var p1 person
	say(d1)
	say(p1)
}

type speaker interface {
	speak() //只要实现了speak的方法的变量都是speaker的类型，指针接受者只能接受指针接受者，值接受者都可以接受
}

type dog struct{}

func (d dog) speak() {
	fmt.Println("我是狗的方法--")
}
func (p person) speak() {
	fmt.Println("我是人的方法--")
}

type person struct{}

func say(s speaker) {
	s.speak()
}
