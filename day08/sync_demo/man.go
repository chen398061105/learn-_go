package main

//互斥锁，只能保证一个goroutine 先执行
//读写互斥锁，
//当一个线程获取读锁之后，其他的oroutine如果是读锁就会继续获得锁，如果是写锁就会等待
//当一个线程获取写锁之后，其他的oroutine都会等待，运用场景：读的次数远远大于写的时候
import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

var x = 0
var wg sync.WaitGroup
var lock sync.Mutex
var rwlock sync.RWMutex

func main() {
	// wg.Add(2)
	// go add()
	// go add()
	// wg.Wait()
	// fmt.Println(x)
	//读写测试
	// start := time.Now()
	// for i := 0; i < 10; i++ {
	// 	go write()
	// 	wg.Add(1)
	// }
	// for i := 0; i < 1000; i++ {
	// 	go read()
	// 	wg.Add(1)
	// }
	// wg.Wait()
	// fmt.Println(time.Now().Sub(start))
	doMap()

}

//读写互斥锁案例
func read() {
	defer wg.Done()
	// lock.Lock()
	rwlock.Lock()
	fmt.Println(x)
	time.Sleep(time.Millisecond)
	// lock.Unlock()
	rwlock.Unlock()
}
func write() {
	defer wg.Done()
	// lock.Lock()
	rwlock.Lock()
	x = x + 1
	time.Sleep(time.Millisecond * 5)
	// lock.Unlock()
	rwlock.Unlock()
}

// //互斥锁案例
// func add() {
// 	defer wg.Done()
// 	for i := 0; i < 50000; i++ {
// 		lock.Lock() //使用互斥锁能够保证同一时间有且只有一个goroutine进入临界区，其他的等待，
// 		x = x + 1
// 		lock.Unlock()
// 	}
// }
//并发内置map
var m = sync.Map{}

func doMap() {
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			m.Store(key, n)
			value, _ := m.Load(key)
			fmt.Printf("K:%v,value:%v\n", key, value)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

//读写互斥锁案例
