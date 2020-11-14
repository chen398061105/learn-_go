package main

import (
	"fmt"
	"sync"
)

//var 变量 chan 元素类型 先进先出的原则
//var cha1 chan int 申明一个cha1的管道
var ch1 chan int
var ch2 chan int
var wg sync.WaitGroup

func main() {
	ch1 = make(chan int, 50)
	ch2 = make(chan int, 100)
	wg.Add(2)
	go demo1()
	go demo2()
	wg.Wait()
	for ret := range ch2 {
		fmt.Println(ret)
	}
}
func demo1() {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		ch1 <- i
	}
	close(ch1)
}
func demo2() {
	wg.Done()
	for {
		x, ok := <-ch1
		if !ok {
			break
		}
		ch2 <- x * x
	}
	close(ch2)
}
