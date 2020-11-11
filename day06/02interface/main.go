package main

import "fmt"

//汽车接口类型的联系
func main() {
	var f1 = falali{
		name: "法拉利",
	}
	var b1 = baoshijie{
		name: "保时捷",
	}

	drive(f1)
	drive(b1)
}

//通用car方法接口
type car interface {
	run()
}

func drive(c car) {
	c.run()
}

type baoshijie struct {
	name string
}
type falali struct {
	name string
}

func (b baoshijie) run() {
	fmt.Printf("车牌名%s 跑", b.name)
}
func (f falali) run() {
	fmt.Printf("车牌名%s 跑", f.name)
}
