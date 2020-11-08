package main

import "fmt"

//函数为的语句必须是关键字开头
// var age int
// age =1
func main() {
	// age := 18 函数内部专属定义变量方法
	//外部的话  var name string  或者  var name = "hello"

	// if age > 18 {
	// 	fmt.Println("不大于18")
	// } else {
	// 	fmt.Println("大与18")
	// }
	//判断时候直接赋值，减少内存，作用域只在当前判断内
	// if num := 10; num > 1 {
	// 	fmt.Println("num > 10")
	// }
	//基本格式
	// for i := 0; i < 5; i++ {
	// 	fmt.Println(i)
	// }
	//for rang循环
	// s := "hello"
	// for index, value := range s {
	// 	fmt.Printf("%d %c \n", index, value)
	// }
	// 9*9乘法表
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%v * %v = %v\t", j, i, j*i)
		}
		fmt.Println("")
	}

}
