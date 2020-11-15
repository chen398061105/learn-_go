package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

//udp 客户端
func main() {
	socket, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 10000,
	})
	if err != nil {
		fmt.Printf("udp failed,err:%v", err)
		return
	}
	defer socket.Close()
	reader := bufio.NewReader(os.Stdin)
	var replay [1024]byte
	for {
		fmt.Print("请输入内容：")
		msg, _ := reader.ReadString('\n')
		socket.Write([]byte(msg))
		//收到回复的数据
		n, _, err := socket.ReadFromUDP(replay[:])
		if err != nil {
			fmt.Printf("socket failed,err:%v", err)
			return
		}
		fmt.Println("收到回复:", string(replay[:n]))
	}
}
