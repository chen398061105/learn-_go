package main

import "fmt"

func main() {
	p := person{
		name: "xiaoming",
		age:  12,
		address: address{
			province: "福州",
			city:     "福清",
		},
	}
	fmt.Println(p.address.city)
	fmt.Println(p.city)
}

type person struct {
	name string
	age  int64
	address
	// address address 嵌套结构体
}

type address struct {
	province string
	city     string
}
