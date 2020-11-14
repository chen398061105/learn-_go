package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	// f()
	for i := 0; i < 10; i++ {
		wg.Add(1) //执行顺序 1 启动循环时，每次接收器wg +1
		go f1(i)
	}
	//如何知道goroutine都结束了
	//引入这个 ync.WaitGroup
	wg.Wait() //执行顺序3

}
func f() {
	rand.Seed(time.Now().UnixNano()) //随机种子
	for i := 0; i < 5; i++ {
		r1 := rand.Int()
		r2 := rand.Intn(10) // 0 <=x <10
		fmt.Println(r1, r2)
	}
}
func f1(i int) {
	defer wg.Done() //执行顺序2
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(300)))
	fmt.Println(i)
}

/*GMP goroutine调度模型
G 就是goroutine 初始栈 大小2k
M go运行队操作系统内的线程的虚拟，M与内线程一般是一一映射关系，goroutine最终要放在m上执行
P 管理一组G队列
*/
