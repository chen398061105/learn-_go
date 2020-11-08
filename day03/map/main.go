package main

import "fmt"

func main() {
	// var m1 map[string]int         //定义一个key为strng 值为int的map
	// m1 = make(map[string]int, 10) //设置内存时候，尽量一开始就设置好
	// m1["hello"] = 18
	// m1["world"] = 20
	// fmt.Println(m1)
	// //约定俗成 ok 布尔值
	// value, ok := m1["a"]
	// if !ok {
	// 	fmt.Print("不存在该key")
	// } else {
	// 	fmt.Println(value)
	// }

	// delete(m1, "hello") //如果nil 不操作
	// fmt.Println(m1)

	//元素类型为map的切片
	var s1 = make([]map[int]string, 1, 1)
	s1[0] = make(map[int]string) //切片和map都要 make初始化
	s1[0][1] = "he"              //key 1 value he
	fmt.Println(s1)

	//值类型为切片的map
	var m1 = make(map[int][]int, 1)
	m1[1] = []int{1, 2, 3}
	fmt.Println(m1)

}
