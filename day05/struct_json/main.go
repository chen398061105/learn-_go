package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	p := person{
		Name: "小明",
		Age:  18,
	}
	//序列化
	p1, err := json.Marshal(p)
	if err != nil {
		fmt.Printf("%v ", err)
		return
	}
	fmt.Printf("%#v\n", string(p1))
	//反序列化
	str := `{"name":"小明","age":11}`
	// json.Unmarshal([]byte(str), p)
	json.Unmarshal([]byte(str), &p) //如果不传指针不会修改值
	fmt.Printf("%#v\n", p)
}

type person struct {
	Name string `json:"name"`
	Age  int64  `json:"age"`
}
