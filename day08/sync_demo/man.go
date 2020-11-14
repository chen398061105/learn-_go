package main

//互斥锁，只能保证一个goroutine 先执行
//读写互斥锁，
//当一个线程获取读锁之后，其他的oroutine如果是读锁就会继续获得锁，如果是写锁就会等待
//当一个线程获取写锁之后，其他的oroutine都会等待，运用场景：读的次数远远大于写的时候
import (
	"fmt"
	"sync"
)

var x = 0
var wg sync.WaitGroup
var lock sync.Mutex
var rwlock sync.RWMutex

func main() {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)
}

//互斥锁案例
func add() {
	defer wg.Done()
	for i := 0; i < 50000; i++ {
		lock.Lock() //使用互斥锁能够保证同一时间有且只有一个goroutine进入临界区，其他的等待，
		x = x + 1
		lock.Unlock()
	}
}

//读写互斥锁案例
