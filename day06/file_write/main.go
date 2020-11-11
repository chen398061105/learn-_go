package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// writeFile()
	// writeFile2()
	// writeFile3()
	test()
}
func writeFile() {
	fileObj, err := os.OpenFile("./xxx.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0775)
	if err != nil {
		fmt.Printf("open file failed .err:%v", err)
		return
	}
	defer fileObj.Close()
	//字节写入
	fileObj.Write([]byte("aaaa bbb ccc"))
	//
	fileObj.WriteString("你好啊")
}
func writeFile2() {
	fileObj, err := os.OpenFile("./xxx.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0775)
	if err != nil {
		fmt.Printf("open file failed .err:%v", err)
		return
	}
	defer fileObj.Close()
	wr := bufio.NewWriter(fileObj)
	wr.WriteString("你好2")
	wr.Flush()
}
func writeFile3() {
	str := "good"
	err := ioutil.WriteFile("./xxx.txt", []byte(str), 0775)
	if err != nil {
		return
	}
}

func test() {
	var input string
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("请输入与内容：")
	input, err := reader.ReadString('\n')
	if err != nil {
		return
	}
	fmt.Printf("你输入的内容是：%s", input)
}
