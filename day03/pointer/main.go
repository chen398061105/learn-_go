package main

import "fmt"

// & 取地址  * 根据地址取值
func main() {
	n := 18
	p := &n
	fmt.Printf("%T", p)
	m := *p
	fmt.Println(m)
}

//make 和 new的区别
//new 给基本数据类型string int申请内存，返回的对应的本身的指针类型
//make 给map，chan，slice 申请内存 返回的是各自对应类型的本身
