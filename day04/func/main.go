package main

import "fmt"

//defer 函数执行后再执行
func main() {
	// fmt.Println(1)
	// defer fmt.Println(2)
	// fmt.Println(3)
	f1 := func() {
		fmt.Println("我是匿名函数1")
	}
	f1()

	f2 := func(x, y int) {
		fmt.Println("我是匿名函数2", x+y)
	}
	f2(1, 2)

	//初始化，只调用一次，即立即执行行数
	func(x, y int) {
		fmt.Println("我是匿名函数3", x+y)
	}(100, 200)

}

//匿名函数一般在函数内部使用，因为函数内部不能再次申明带有名字的函数
