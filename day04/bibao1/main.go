package main

import (
	"fmt"
	"strings"
)

//判断是否含有后缀名 闭包练习
func main() {
	jpgTest := makeSuffix(".jpg")
	txtTest := makeSuffix(".txt")

	fmt.Println(jpgTest("test.jpg"))
	fmt.Println(txtTest("test"))

}
func makeSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}
