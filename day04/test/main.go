package main

import (
	"fmt"
	"strings"
)

//计算单词出现次数
func main() {
	str := "how do you do"
	//空格切割字符串,返回切片
	s1 := strings.Split(str, " ")
	fmt.Printf("%T", s1)
	//存到map里面
	m1 := make(map[string]int, 10)
	//遍历map，如果存在则1不存在+1
	for _, v := range s1 {
		// fmt.Println(v)
		if _, ok := m1[v]; !ok {
			m1[v] = 1
		} else {
			m1[v]++
		}
	}
	for k, v := range m1 {
		fmt.Println(k, v)
	}
	fmt.Println("-----回文判断-------")
	str1 := "a山西运煤车煤运西山a"
	r := make([]rune, 0, len(str1))
	for _, v := range str1 {
		r = append(r, v)
	}
	for i := 0; i < len(r)/2; i++ {
		if r[i] != r[len(r)-1-i] {
			fmt.Println("不是回文")
			return
		}
	}
	fmt.Println("是回文")
}
