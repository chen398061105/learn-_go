package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var flag bool //通过设置标签退出

func f() {
	defer wg.Done()
	for {
		fmt.Println("test")
		time.Sleep(time.Millisecond * 500)
		if flag {
			break
		}
	}
}

func main() {
	wg.Add(1)
	go f()
	time.Sleep(time.Second * 5)
	flag = true
	wg.Wait()
	//如何优雅的退出goroutine
}
