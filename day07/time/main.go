package main

import (
	"fmt"
	"time"
)

func main() {
	//获取当前时间
	now := time.Now()
	fmt.Println(now)
	//时间戳
	fmt.Println(now.Unix())     //秒
	fmt.Println(now.UnixNano()) //毫秒
	//转换时间格式
	ret := time.Unix(1605161928, 0)
	fmt.Println(ret)
	//当前时间加24小时
	fmt.Println(now.Add(24 * time.Hour))
	//定时器
	// time := time.Tick(time.Second)
	// for t := range time {
	// 	fmt.Println("每隔一秒输出了哦", t)
	// }
	//时间格式化  go诞生时间 2006年1月2号15点04分05秒 2006 12345
	fmt.Println(now.Format("2006-01-02 15:04:05"))
	//当前时间 am
	fmt.Println(now.Format("2006-01-02 03:04:05 PM"))
	//sleep
	n := 3
	fmt.Println("计时开始")
	time.Sleep(time.Duration(n) * time.Second)
	// time.Sleep(5 * time.Second) 和上面的一样
	fmt.Println("计时结束")

}
