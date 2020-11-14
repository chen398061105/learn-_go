package main

import (
	"fmt"
	"sync"
)

//var 变量 chan 元素类型 先进先出的原则
//var cha1 chan int 申明一个cha1的管道
var b chan int
var wg sync.WaitGroup

/*通道的操作
发送 ch <- 10 把10发送到ch中
接收 x:= <- ch 从ch接收值并赋值给x  <-ch 从ch中接收值并忽略结果
关闭 close(ch)
*/
func main() {
	bufferChannel()
}
func noBufferChannel() {
	b = make(chan int) //通道的初始化
	wg.Add(1)          // 1 相当于 1个go
	go func() {
		defer wg.Done()
		x := <-b
		fmt.Println("后台接收到x了", x)
	}()
	b <- 10 //不带缓冲区的，如果没有上面go，就会死机 卡主
	fmt.Println("我要吧b发送到通道去了哦")
	fmt.Println(b)
	wg.Wait()
}

//带缓冲区
func bufferChannel() {
	b = make(chan int, 1) //通道的初始化
	b <- 10
	fmt.Println("我要吧b发送到通道去了哦")
	fmt.Println(b)
}
