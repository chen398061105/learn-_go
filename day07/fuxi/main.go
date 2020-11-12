package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	// open1()
	// open2()
	// open3()
	insertInfo()
}

//中间插入数据
func insertInfo() {

}
func open1() {
	file, err := os.Open("./main.go")
	// defer file.Close() 当err有值的时候，file就是nill，nill不能调用close方法
	if err == io.EOF {
		return
	}
	if err != nil {
		fmt.Printf("err:%v", err)
		return
	}
	defer file.Close()
	//缓冲区
	reader := bufio.NewReader(file)
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
func open2() {
	file, err := os.Open("./main.go")
	if err != nil {
		fmt.Printf("err:%v", err)
		return
	}
	defer file.Close()
	var tmp = make([]byte, 128)
	for {
		n, err := file.Read(tmp)
		if err == io.EOF {
			fmt.Printf("err :%v", err)
			return
		}
		if err != nil {
			fmt.Printf("err:%v", err)
			return
		}
		fmt.Print(string(tmp[:n]))
		if n < 128 {
			return
		}
	}
}
func open3() {
	file, err := ioutil.ReadFile("./main.go")
	if err != nil {
		fmt.Printf("err:%v", err)
		return
	}
	fmt.Println(string(file))
}
