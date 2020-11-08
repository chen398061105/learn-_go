package main

import "fmt"

//func 函数名（参数）（返回值）{函数体}
func main() {
	res := sum(1, 2)
	fmt.Println(res)

	_, r1 := sum1()
	fmt.Println(r1)

	fun1(1, 2, 3, 4)

}

// func sum(x int, y int)  int {
// 	return x + y
// } 匿名返回值 ，下面的相当于声明一个result 的返回
func sum(x int, y int) (result int) {
	result = x + y
	return //result 可写可不写
}

//参数类型一致时候，前面的参数类型可以省略简写
// func sum(x , y int)  int {
// 	return x + y
// }
func sum1() (int, string) {
	return 1, "he"
}

func fun1(x ...int) {
	fmt.Println(x) //返回的是一个切片
}
