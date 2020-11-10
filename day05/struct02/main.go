package main

import "fmt"

//结构体是值类型  type
type person struct {
	name string //字段名 + 字段类型 不能重复
}

// go 里面参数传参永远是 粘贴复制 引用类型， 要改变值 需要加*
var f = func(x *person) {
	// (*x).name = "aa"  和下面一样，自动推测
	x.name = "aa"
}

func main() {
	var chen person
	chen.name = "小明"
	f(&chen)
	fmt.Println(chen.name)
}
