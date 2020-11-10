package main

import "fmt"

//结构体是值类型  type
type person struct {
	name   string //字段名 + 字段类型 不能重复
	age    int
	gender string
	hobby  []string
}

func main() {
	var chen person
	chen.name = "小明"
	fmt.Println(chen.name)
	//匿名结构体,多用于临时场景
	var s struct {
		x int
		y int
	}
	s.x = 1
	s.y = 2
	fmt.Println(s)
}
