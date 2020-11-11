package main

import (
	"fmt"
	"os"
)

func showMenu() {
	fmt.Println("===学生管理系统===")
	fmt.Print("1 查看\t2 添加\t3 编辑\t 4 删除\t5 退出\n请选择：")
}

func main() {
	smr := studentInfo{
		allStudent: make(map[int64]student, 50),
	}
	for {
		showMenu()
		var input int
		fmt.Scan(&input)
		fmt.Printf("您选择了%d 选项\n", input)

		switch input {
		case 1:
			smr.show()
		case 2:
			smr.add()
		case 3:
			smr.edit()
		case 4:
			smr.del()
		case 5:
			os.Exit(0)
			fmt.Print("退出成功\n")
		default:
			fmt.Printf("无效的选择，请重新输入")
		}
	}
}
