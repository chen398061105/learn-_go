package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

//服务端
//处理函数
func process(conn net.Conn) {
	defer conn.Close()
	getClientinfo := bufio.NewReader(os.Stdin)
	var buf [1024]byte
	for {
		reader := bufio.NewReader(conn)
		n, err := reader.Read(buf[:]) //读取数据
		if err != nil {
			fmt.Printf("read from client failed,err:%v", err)
			break
		}
		recStr := string(buf[:n])
		fmt.Println("收到客户端信息：", recStr)

		//回复客户端信息
		fmt.Print("请回复：")
		msg, _ := getClientinfo.ReadString('\n')
		msg = strings.TrimSpace(msg)
		if msg == "exit" {
			break
		}
		conn.Write([]byte(msg)) //发送数据
	}
}
func main() {
	//本地端口服务
	listen, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Printf("listen failed,err:%v", err)
		return
	}
	defer listen.Close()
	//等待别人来跟我建立连接
	for {
		conn, err := listen.Accept() //建立连接
		if err != nil {
			fmt.Printf("Accept failed,err:%v", err)
			continue
		}
		//与客户端通信
		go process(conn) //启动一个goroutine处理连接
	}
}
