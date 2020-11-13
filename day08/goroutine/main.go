package main

import (
	"fmt"
	"time"
)

func main() {
	go hello() //单独开启一个goroutine去执行hello函数
	//开始休眠，可以执行上面的
	time.Sleep(time.Second)
	fmt.Println("main")
	//main结束了，有main函数执行的goroutine也都结束了
}
func hello() {
	fmt.Println("hello")
}
