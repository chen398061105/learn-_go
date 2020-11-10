package main

import "fmt"

func main() {
	animal := newDog("哈士奇")
	animal.wang()
}

//结构体
type dog struct {
	name string
}

//构造函数
func newDog(name string) dog {
	return dog{
		name: name,
	}
}

//方法是作用特点类型的函数，接受者表示调用该方法具体变量。 约定 多用类型的首字母的小写 d dog
func (d dog) wang() {
	fmt.Printf("%s wang wang", d.name)
}
