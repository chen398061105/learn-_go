package main

import (
	"fmt"
	"strconv"
)

func main() {

	// ret, err := strconv.ParseInt(str, 10, 32)
	// if err != nil {
	// 	return
	// }
	//字符串转数字
	str := "1000"
	ret, err := strconv.Atoi(str) //推荐
	if err != nil {
		return
	}
	fmt.Printf("type:%T value:%v\n", ret, ret)
	//数字转字符串
	nun := 10
	str1 := strconv.Itoa(nun)
	fmt.Printf("type:%T value:%v\n", str1, str1)
	//字符串转布尔
	boolValue, _ := strconv.ParseBool("true")
	fmt.Printf("type:%T value:%v\n", boolValue, boolValue)

}
