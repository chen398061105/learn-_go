package main

import "fmt"

func main() {
	ret := add()
	// ret := add(100) 闭包函数，包含外部作用域的变量， 可以做为返回值，变量先自己后外层
	ret1 := ret(100)
	fmt.Println(ret1)
}

func add() func(int) int {
	var x = 10
	return func(y int) int {
		x += y
		return x
	}
}

// func add() func(int) int {
// 	return func(y int) int {
// 		x += y
// 		return x
// 	}
// }
