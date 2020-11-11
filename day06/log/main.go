package main

import (
	"fmt"
	"os"
)

//日志收集 日志级别
//往不同地方写入日志
//logger.Trace() logger.Debug() logger.Wadnig() logger.Info() logger.Error()
func main() {
	fmt.Fprintln(os.Stdout, "这是一条日志")
	fileObj, _ := os.OpenFile("./xxxx.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0775)
	fmt.Fprintln(fileObj, "这是一条日志")
}
