package main

import "fmt"

func main() {
	d := dog{
		animal: animal{name: "xiaobai"},
		feet:   4,
	}
	d.wang()
	d.move()
}

type animal struct {
	name string
}

func (a animal) move() {
	fmt.Printf("%s 在 移动", a.name)
}

type dog struct {
	feet   uint8
	animal //动物拥有的方法，狗都拥有
}

func (d dog) wang() {
	fmt.Printf("%s 在 汪汪", d.name)
}
