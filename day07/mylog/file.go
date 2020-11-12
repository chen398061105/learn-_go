package mylog

import (
	"fmt"
	"os"
	"path"
	"time"
)

// FileLogger 文件记录日志
type FileLogger struct {
	level       LogLevel // 日志错误等级
	filePath    string   //文件保存路径
	fileName    string   //文件保存名
	maxFileSize int64    // 文件最大保存容量
	fileObj     *os.File
	errFileObj  *os.File
}

// NewFileLogger 构造函数
func NewFileLogger(leverStr, filePath, fileName string, maxSize int64) *FileLogger {
	logLevel, err := parseLogLevel(leverStr)
	if err != nil {
		panic(err)
	}
	fl := &FileLogger{
		level:       logLevel,
		filePath:    filePath,
		fileName:    fileName,
		maxFileSize: maxSize,
	}
	//初始化 打开文件路径和文件名
	err = fl.initFile()
	if err != nil {
		panic(err)
	}
	return fl
}
func (f *FileLogger) initFile() error {
	//普通日志
	fulelFileName := path.Join(f.filePath, f.fileName)
	fileObj, err := os.OpenFile(fulelFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open log file failed,err:%v\n", err)
		return err
	}
	//专门的错误日志
	errFulelFileName := path.Join(f.filePath, f.fileName)
	errFileObj, err := os.OpenFile("err_"+errFulelFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open err log file failed,err:%v\n", err)
		return err
	}
	//日志文件都打开完毕
	f.fileObj = fileObj
	f.errFileObj = errFileObj
	return nil
}
func (f *FileLogger) enable(LogLevel LogLevel) bool {
	return LogLevel >= f.level
}

//按照文件大小切割
func (f *FileLogger) checkSize(file *os.File) bool {
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("get file info failed,err:%v", err)
		return false
	}
	//如果当前文件大小大于等于 传入的文件大小则返回true
	return fileInfo.Size() >= f.maxFileSize
}

//切割文件函数
func (f *FileLogger) splitFile(file *os.File) (*os.File, error) {
	//需要切割的文件
	now := time.Now().Format("2006-01-02 15:04:05")
	//获取文件信息
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("get file info failed,err:%v", err)
		return nil, err
	}
	//备份新文件 xx.log->xxx.log.bak20201111
	oldFile := path.Join(f.filePath, fileInfo.Name())
	newFileName := fmt.Sprintf("%s.bak%s", oldFile, now)
	//关闭当前文件
	file.Close()
	os.Rename(oldFile, newFileName)
	//打开新的文件
	fileObj, err := os.OpenFile(oldFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open new file failed,err:%v", err)
		return nil, err
	}
	//将打开的新文件对象打开f.fileObj
	return fileObj, nil
}

//  文件 输出格式
func (f *FileLogger) logFormat(level LogLevel, fortmat string, a ...interface{}) {
	if f.enable(level) {
		now := time.Now()
		time := now.Format("2006-01-02 15:04:05")
		funcName, fileName, lineNo := getNoInfo(2) //2 是个隔着的层数，main入口开始，函数调用开始算一层
		msg := fmt.Sprintf(fortmat, a...)
		// 既可以传字符串，也可以传其他的报错类型,大小判断
		if f.checkSize(f.fileObj) {
			newFile, err := f.splitFile(f.fileObj)
			if err != nil {
				return
			}
			f.fileObj = newFile
		}
		fmt.Fprintf(f.fileObj, "[%s][%s][funcName:%s][fileName:%s][Line:%d]%s \n", time, getLogLevelString(level), funcName, fileName, lineNo, msg)
		//如果要记录的日志大与error级别，另外单独在记录一份
		if level >= ERROR {
			if f.checkSize(f.errFileObj) {
				newFile, err := f.splitFile(f.errFileObj)
				if err != nil {
					return
				}
				f.errFileObj = newFile
			}
			fmt.Fprintf(f.errFileObj, "[%s][%s][funcName:%s][fileName:%s][Line:%d]%s \n", time, getLogLevelString(level), funcName, fileName, lineNo, msg)
		}
	}
}

// Close 关闭文件
func (f *FileLogger) Close() {
	f.fileObj.Close()
	f.errFileObj.Close()
}

// Debug 调试
func (f *FileLogger) Debug(format string, a ...interface{}) {
	f.logFormat(DEBUG, format, a...)
}

// Trace 追踪
func (f *FileLogger) Trace(format string, a ...interface{}) {
	f.logFormat(TRACE, format, a...)
}

// Info 正常信息
func (f *FileLogger) Info(format string, a ...interface{}) {
	f.logFormat(INFO, format, a...)
}

// Warning 警告
func (f *FileLogger) Warning(format string, a ...interface{}) {
	f.logFormat(WARNING, format, a...)
}

// Error 错误
func (f *FileLogger) Error(format string, a ...interface{}) {
	f.logFormat(ERROR, format, a...)
}

// Fatal 严重错误
func (f *FileLogger) Fatal(format string, a ...interface{}) {
	f.logFormat(FATAL, format, a...)
}
