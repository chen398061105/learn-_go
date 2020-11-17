package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func f(cxt context.Context) {
	defer wg.Done()
	go f2(cxt)
LOOP:
	for {
		fmt.Println("test")
		time.Sleep(time.Millisecond * 500)
		select {
		case <-cxt.Done(): //等待通知
			break LOOP
		default:
		}
	}
}
func f2(cxt context.Context) {
	defer wg.Done()
LOOP:
	for {
		fmt.Println("test-son")
		time.Sleep(time.Millisecond * 500)
		select {
		case <-cxt.Done(): //等待通知
			break LOOP
		default:
		}
	}
}

func main() {
	cxt, cancel := context.WithCancel(context.Background()) //参数 父节点 返回没有参数的没有返回值的cancel
	wg.Add(1)
	go f(cxt)
	time.Sleep(time.Second * 5)
	//如何优雅的退出goroutine
	cancel()
	wg.Wait()

}
