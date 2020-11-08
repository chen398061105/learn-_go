package main

import "fmt"

//常量
const pi = 3.1415

//批量申明多行时候，后面如果不写则和前面一页
const (
	statusOK = 100
	statusNot
)

//iota 每次加一 ,常量中每新增一行 加一次计算1
const (
	n1 = iota //0
	n2        //1
	n3        // 2
	/*
		c1 = iota  0
		c2 =100    100
		c3 =itoa   2
		c4         3
	*/
)

func main() {
	var (
		name string
		age  int
	) //全局函数内都可以 常用
	name = "xiaoming"
	age = 188
	fmt.Printf("name :%s age : %v", name, age)
	var s1 = "s1" //类推导声明
	s2 := "s1"    //简短变量申明，只能在函数里面用 常用
	fmt.Println(s1)
	fmt.Print(s2)

	fmt.Println(pi)
	fmt.Println("Ok", statusOK)
	fmt.Println("not", statusNot)
	fmt.Println("n1", n1)
	fmt.Println("n2", n2)
	fmt.Println("n3", n3)

}
