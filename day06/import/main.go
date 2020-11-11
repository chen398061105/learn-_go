package main

import (
	"fmt"

	tool "github.com/chen398061105/study/tool"
	//src目录开始，如果要引用其他包里面的标识符（变量，常量，类型，函数等） 该标识符首字母必须大写
)

func main() {
	ret := tool.Add(1, 2)
	fmt.Println(ret)
}
