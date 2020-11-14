package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//使用goroutine和channel计算一个int64的随机数和的程序
//1 开启一个goroutine循环生成int64的随机数，发送到jobchan
//2 开启24个goroutine从jobChan中取出随机数计算各位数的和，发送到resultChan中
//3 主goroutine从resultChan中取出结果并打印到终端
type job struct {
	value int64
}

type result struct {
	job    *job
	result int64
}

var jobChan = make(chan *job, 100)
var resultChan = make(chan *result, 100)
var wg sync.WaitGroup

func worker(ch1 chan<- *job) {
	defer wg.Done()
	for { //随机生成int64的随机数 存到 ch1
		x := rand.Int63()
		newJob := &job{
			value: x,
		}
		ch1 <- newJob
		time.Sleep(time.Second)
	}
}
func getResult(ch1 <-chan *job, resultChan chan<- *result) {
	defer wg.Done()
	//开启24个goroutine从jobChan中取出随机数计算各位数的和，发送到resultChan中
	for {
		//ch1 值取出来
		job := <-ch1
		//计算各位数的和
		sum := int64(0)
		n := job.value
		for n > 10 {
			sum = +n % 10
			n = n / 10
		}
		newResult := &result{
			job:    job,
			result: sum,
		}
		resultChan <- newResult
	}

}
func main() {
	wg.Add(1)
	go worker(jobChan)
	wg.Add(24)
	for i := 0; i < 24; i++ {
		go getResult(jobChan, resultChan)
	}
	for res := range resultChan {
		fmt.Printf("value:%v sum:%d\n", res.job.value, res.result)
	}
	wg.Wait()
}
