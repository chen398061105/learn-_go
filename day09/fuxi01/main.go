/*1 互斥锁 sync.Mutex 底层是一个结构体 值类型 传指针加& 方法：lock 和unlock
    为什么要用锁？防止同一时间多个goroutine操作同一个资源，
  2 读写互斥锁，适用读多写少的场景+
    读的 goroutine 来了获取的是读锁，后续的goroutine能读不能写
    写的 goroutine 来了的是写锁，后续的 不管是什么都要等待获取锁
    方法 var rwLock sync.RWMutex rwLock.Rlock/RUnlock 获取写/释放读写锁   rwLock.lock/unlokc 获取血锁/释放写锁
  3 等待组 var wg sync.Waitgroup 用来等goroutine结束完再继续
   wg.Add(n) n 是对应的goroutine数 wg.done 在对应goroutine函数里面，函数要结束的时候完成技术-1 wg.wait 等待所有goroutine都结束之后才执行后面代码
   *var once sync.Once    once.Do(func(){}闭包函数)只执行一次，成功则true
   4 sync.Map 并发操作一个map的时候，内置的map不是安全的，用系统内置提供的开箱即用
	var syncMap sync.Map  方法 synaMap.Store(key,value)  synaMap.Load(key) 取值
   5 原子操作 var i int64 = 10   atomic.AddInt64(&i,1)
*/
package main
