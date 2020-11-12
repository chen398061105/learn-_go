package mylog

import (
	"errors"
	"fmt"
	"path"
	"runtime"
	"strings"
)

//LogLevel 定义一个日志级别,开关控制用
type LogLevel int16

// 日志级别
const (
	UNKNOW  LogLevel = iota //0
	DEBUG                   //1
	TRACE                   //2
	INFO                    //3
	WARNING                 //4
	ERROR                   //5
	FATAL                   //6
)

//控制传入日志的级别
func parseLogLevel(s string) (LogLevel, error) {
	s = strings.ToLower(s)
	switch s {
	case "debug":
		return DEBUG, nil
	case "trace":
		return TRACE, nil
	case "info":
		return INFO, nil
	case "warning":
		return WARNING, nil
	case "error":
		return ERROR, nil
	case "fatal":
		return FATAL, nil
	default:
		err := errors.New("Invalid log level")
		return UNKNOW, err
	}
}
func getLogLevelString(lv LogLevel) string {
	switch lv {
	case DEBUG:
		return "DEBUG"
	case TRACE:
		return "TRACE"
	case INFO:
		return "INFO"
	case WARNING:
		return "WARNING"
	case ERROR:
		return "ERROR"
	case FATAL:
		return "FATAL"
	default:
		return "DEBUG"
	}
}

//获取行号
func getNoInfo(num int) (funcName, fileName string, lineNo int) {
	//pc 调用函数的相关信息 , file 当前执行的文件名, line 行号, ok true or false
	pc, file, lineNo, ok := runtime.Caller(num)
	if !ok {
		fmt.Printf("runtime failed\n")
		return
	}
	funcName = runtime.FuncForPC(pc).Name()
	fileName = path.Base(file)
	return
}
