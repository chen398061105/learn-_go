package main

import (
	"time"

	myloger "github.com/chen398061105/study/day07/mylog"
)

//日志收集 日志级别 日志开关
//完整的日志要有时间，行号，文件名，日志级别，日志信息，日志切割
//往不同地方写入日志
// logger.Debug() logger.warning() logger.Info() logger.Error() logger.Fatal()严重错误
func main() {
	// log := myloger.NewLog("debug") //控制台
	log := myloger.NewFileLogger("debug", "./", "demo.log", 1024)
	//这是一条Debug日志信息,ID:%d Name:%s", id, name
	//为了可以传其他的东西进去，不只是字符串
	id := 1
	name := "test"
	for {
		log.Debug("这是一条Debug日志信息,ID:%d Name:%s", id, name)
		log.Trace("这是一条Trace日志信息")
		log.Info("这是一条Info日志信息")
		log.Warning("这是一条Warning日志信息")
		log.Error("这是一条Error日志信息")
		log.Fatal("这是一条Fatal日志信息")
		time.Sleep(time.Second)
	}
}
