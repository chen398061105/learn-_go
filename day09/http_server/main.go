package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func f1(w http.ResponseWriter, r *http.Request) {
	html, err := ioutil.ReadFile("./test.html") //相当于网站内容
	if err != nil {
		w.Write([]byte(fmt.Sprintf("err:%v", err)))
	}
	w.Write([]byte(html)) //获取到内容写入到网站，网址是 127.0.0.1:9000/test/
}
func f2(w http.ResponseWriter, r *http.Request) {
	//fmt.Println(r.URL) //对于get的请求，参数都放在url上面，请求体是没有
	//自动获取 query
	queryName := r.URL.Query().Get("name")
	fmt.Println("name", queryName)
	fmt.Println("方法", r.Method)
	fmt.Println(ioutil.ReadAll(r.Body))
	w.Write([]byte("test2"))
}

//http 服务器
func main() {
	http.HandleFunc("/test2/", f2)
	http.HandleFunc("/test/", f1) //test是网站路径，f1是接收的请求
	http.ListenAndServe("127.0.0.1:9000", nil)
}
