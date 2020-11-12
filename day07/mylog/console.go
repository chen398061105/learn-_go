package mylog

import (
	"fmt"
	"time"
)

// ConsoleLogger 日志结构体
type ConsoleLogger struct {
	Level LogLevel //设置级别是为了控制需要输出的日志级别
}

// NewLog 构造函数 ConsoleLogger 返回值
func NewLog(levStr string) ConsoleLogger {
	level, err := parseLogLevel(levStr)
	if err != nil {
		fmt.Printf("err:%v", err)
		panic(err)
	}
	return ConsoleLogger{
		Level: level,
	}
}
func (c ConsoleLogger) enable(LogLevel LogLevel) bool {
	return LogLevel >= c.Level
}

// Debug 调试
func (c ConsoleLogger) Debug(format string, a ...interface{}) {
	c.logFormat(DEBUG, format, a...)
}

// Trace 追踪
func (c ConsoleLogger) Trace(format string, a ...interface{}) {
	c.logFormat(TRACE, format, a...)
}

// Info 正常信息
func (c ConsoleLogger) Info(format string, a ...interface{}) {
	c.logFormat(INFO, format, a...)
}

// Warning 警告
func (c ConsoleLogger) Warning(format string, a ...interface{}) {
	c.logFormat(WARNING, format, a...)
}

// Error 错误
func (c ConsoleLogger) Error(format string, a ...interface{}) {
	c.logFormat(ERROR, format, a...)
}

// Fatal 严重错误
func (c ConsoleLogger) Fatal(format string, a ...interface{}) {
	c.logFormat(FATAL, format, a...)
}

// 控制台输出格式
func (c ConsoleLogger) logFormat(level LogLevel, fortmat string, a ...interface{}) {
	if c.enable(level) {
		now := time.Now()
		time := now.Format("2006-01-02 15:04:05")
		funcName, fileName, lineNo := getNoInfo(3) //2 是个隔着的层数，main入口开始，函数调用开始算一层
		msg := fmt.Sprintf(fortmat, a...)          //既可以传字符串，也可以传其他的报错类型
		fmt.Printf("[%s][%s][funcName:%s][fileName:%s][Line:%d]%s \n", time, getLogLevelString(level), funcName, fileName, lineNo, msg)
	}
}
