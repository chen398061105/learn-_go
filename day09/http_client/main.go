package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

//http 客户端
func main() {
	respone, err := http.Get("http://127.0.0.1:9000/test2/?page=1&name=1")
	if err != nil {
		fmt.Printf("get url err:%v", err)
		return
	}
	defer respone.Body.Close()
	//在客户端取得所有服务端发过来的内容
	b, err := ioutil.ReadAll(respone.Body)
	if err != nil {
		fmt.Printf("read body err:%v", err)
		return
	}
	fmt.Println(string(b))
}
