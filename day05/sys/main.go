package main

import (
	"fmt"
	"os"
)

/*
函数版本学生管理系统
查看、删除，添加
*/
type student struct {
	id   int64
	name string
}

func newPerson(id int64, name string) *student {
	return &student{
		id:   id,
		name: name,
	}
}

var (
	allStudent map[int64]*student
)

func show() {
	for k, v := range allStudent {
		fmt.Printf("学号：%d 姓名：%s \n", k, v.name)
	}
}
func del() {
	var (
		delID int64
	)
	fmt.Print("请输入要删除的学号ID：")
	fmt.Scan(&delID)
	delete(allStudent, delID)
	fmt.Print("删除成功\n")
}
func add() {
	var (
		stuID   int64
		stuName string
	)
	fmt.Print("请输入学号：")
	fmt.Scan(&stuID)
	fmt.Print("请输入姓名：")
	fmt.Scan(&stuName)

	newStu := newPerson(stuID, stuName)
	allStudent[stuID] = newStu
	fmt.Print("添加成功\n")

}

func main() {
	allStudent = make(map[int64]*student, 50)
	for {
		fmt.Println("===学生管理系统===")
		fmt.Print("1 查看\t2 删除\t3 添加\t 4 退出\n请选择：")
		var input int
		fmt.Scan(&input)
		fmt.Printf("您选择了%d 选项\n", input)

		switch input {
		case 1:
			show()
		case 2:
			del()
		case 3:
			add()
		case 4:
			os.Exit(0)
			fmt.Print("退出成功\n")
		default:
			fmt.Printf("无效的选择，请重新输入")
		}
	}
}
