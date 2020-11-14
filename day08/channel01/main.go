package main

import "fmt"

func main() {

}

//单向通道，多用于函数的参数中，限制通道只能做什么内容
//ch1 只能存int值，ch2只能取int值  chan前取后存
func single(ch1 chan<- int, ch2 <-chan int) {

}
func test() {
	ch1 := make(chan int, 2) //2 缓冲值 就是可以添加的次数
	ch1 <- 1
	ch1 <- 2
	close(ch1)
	//通道关闭了 仍然可以取值，但是超出范围后，就会取得对应的类型的0 false之类的值
	x, ok := <-ch1
	fmt.Println(x, ok) //1 true
	x, ok = <-ch1
	fmt.Println(x, ok) //2 true
	x, ok = <-ch1
	fmt.Println(x, ok) //0 false
}
