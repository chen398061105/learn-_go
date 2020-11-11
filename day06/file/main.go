package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

//io读
func main() {
	// readFileByIo()
	readFileByAll()
}
func readFileByIo() {
	//打开文件，返回一个文件的结构体
	fileObj, err := os.Open("./main.go")
	if err != nil {
		fmt.Printf("read file faild %v", err)
	}
	defer fileObj.Close()
	//先创建一个文件来读取文件对象
	reader := bufio.NewReader(fileObj)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Printf("read file faild2 %v", err)
			return
		}
		fmt.Print(line)
	}
}

//一次读取全部
func readFileByAll() {
	//放回的是字节 要转string
	ret, err := ioutil.ReadFile("./main.go")
	if err != nil {
		fmt.Printf("read file faild err:%v", err)
		return
	}
	fmt.Print(string(ret))
}
