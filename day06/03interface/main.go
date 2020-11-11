package main

import "fmt"

//接口嵌套
func main() {
	c1 := cat{name: "xiao"}
	c1.animal()
	//空接口,interface是关键字 interface{}加上括号就是接口，接口可以接受任意类型
	var m map[string]interface{}
	m = make(map[string]interface{}, 10)
	m["name"] = "熊猫"
	m["age"] = 1
	m["flag"] = true
	m["hobby"] = [...]string{"吃", "喝"}
	fmt.Println(m)

}

//空接口判断
// func assign( a interface{}){
// 	s,ok :=a.(string)
// 	if !ok {
// 		fmt.Println("判断错误",s)
// 	}
// }

type mover interface {
	move()
}
type eater interface{ eat() }

type animal interface {
	mover
	eater
}
type cat struct {
	name string
}

func (c cat) move() {
	fmt.Println("猫走猫步")
}
func (c cat) eat() {
	fmt.Println("猫吃鱼")
}

func (c cat) animal() {
	fmt.Printf("%s 在散步", c.name)
	c.eat()
	c.move()
}
