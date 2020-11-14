package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch := make(chan int, 1)
	go send(ch)
	for {
		x, ok := <-ch
		fmt.Println(x, ok) //阻塞4秒
		time.Sleep(time.Second)
	}

}
func send(ch1 chan<- int) {
	for {
		num := rand.Intn(10)
		ch1 <- num
		time.Sleep(5 * time.Second)
	}
}
