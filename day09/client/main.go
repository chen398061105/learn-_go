package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

//客户端
func main() {
	// 与服务端建立连接
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Printf("dial failed,err:%v", err)
		return
	}
	//返送数据
	reader := bufio.NewReader(os.Stdin)
	for { //如果for循环 ，服务端可能会一次回复n多信息，而不是一句一句的获取
		fmt.Print("请说话：")
		msg, _ := reader.ReadString('\n')
		msg = strings.TrimSpace(msg)
		if msg == "exit" {
			break
		}
		conn.Write([]byte(msg))
	}
	conn.Close()
}
