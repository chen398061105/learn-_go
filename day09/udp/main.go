package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	// UDP 服务器  一般用户视频
	conn, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 10000,
	})
	if err != nil {
		fmt.Printf("udp failed,err:%v", err)
		return
	}
	defer conn.Close()
	//不需要建立连接直收发数据
	var data [1024]byte
	for {
		n, addr, err := conn.ReadFromUDP(data[:])
		if err != nil {
			fmt.Printf("ReadFromUDP failed,err:%v", err)
			return
		}
		fmt.Println(data[:n])

		//往addr发送数据
		replay := strings.ToUpper(string(data[:n]))
		conn.WriteToUDP([]byte(replay), addr)

	}

}
