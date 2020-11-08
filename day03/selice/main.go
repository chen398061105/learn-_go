package main

import "fmt"

func main() {
	//不能直接比较，只能和nil（空）比较
	// fmt.Println("切片的练习，引用类型") 引用型,数据保存在底层数组里 会更变实际的值
	// var s1 []int //定义一个int类型的切片。底层就是数组
	// var s2 []string
	// fmt.Println(s1 == nil)
	// s1 = []int{1, 2, 3}
	// s2 = []string{"one", "two"}
	// fmt.Println(s1, s2)

	// s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	// s1 := s[0:4] //左闭右开
	// fmt.Println(s1)
	// // s2 := s[0:]//0开始所有
	// s2 := s[:] //0开始所有
	// fmt.Println(s2)
	// s[0] = 100
	// fmt.Println(s)
	//判断切片是否为空 len（s）==0
	//创建切片 make
	// s1 := make([]string, 5)
	// fmt.Printf("s1=:%v len9(s1):%d cap(s1):%d\n", s1, len(s1), cap(s1))
	// s1 = append(s1, "hello")
	// fmt.Printf("s1=:%v len9(s1):%d cap(s1):%d\n", s1, len(s1), cap(s1))
	// fmt.Println(s1)
	//切片的复制
	// s1 := []int{1, 2}
	// s2 := make([]int, 3, 3)
	// copy(s2, s1) //复制后的值 就是原来的值 如果原来的值变化了，不影响复制后的值
	// s1[0] = 11
	// fmt.Println(s1, s2)
	//切片的删除
	s1 := [...]int{1, 2, 3, 4, 5}
	s2 := s1[:]
	s2 = append(s2[0:2], s2[3:]...)
	fmt.Println(s2)

}
