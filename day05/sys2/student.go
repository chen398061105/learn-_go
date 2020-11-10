/*
函数版本学生管理系统2 加强版
有一个物件 他保存了一些数据 ---结构体的字段
他有如下的功能   --- 结构体的方法
查看、删除，添加
*/
package main

import "fmt"

type student struct {
	id   int64
	name string
}
type studentInfo struct {
	allStudent map[int64]student
}

func (s studentInfo) add() {
	var (
		stuID   int64
		stuName string
	)
	fmt.Print("请输入学号：")
	fmt.Scan(&stuID)
	fmt.Print("请输入姓名：")
	fmt.Scan(&stuName)
	newStu := student{
		id:   stuID,
		name: stuName,
	}
	s.allStudent[newStu.id] = newStu
	fmt.Print("添加成功\n")
}
func (s studentInfo) show() {
	if len(s.allStudent) == 0 {
		fmt.Println("暂无信息")
		return
	}
	for _, v := range s.allStudent {
		fmt.Println("===学生信息一览===")
		fmt.Printf("学号：%d 姓名：%s \n", v.id, v.name)
	}
}
func (s studentInfo) del() {
	//获取用户输入id
	var stuID int64
	fmt.Print("请输入要删除学号ID：")
	fmt.Scan(&stuID)
	_, ok := s.allStudent[stuID]
	if !ok {
		fmt.Printf("ID为：%d的用户不存在\n", stuID)
		return
	}
	delete(s.allStudent, stuID)
	fmt.Println("删除成功！")
}
func (s studentInfo) edit() {
	//获取用户输入id
	var stuID int64
	fmt.Print("请输入学号：")
	fmt.Scan(&stuID)
	//获取id,判断是否存在
	stuObj, ok := s.allStudent[stuID]
	if !ok {
		fmt.Printf("ID为：%d的用户不存在\n", stuID)
		return
	}
	fmt.Printf("当前学生姓名为:%s \t", stuObj.name)
	//获取新的名字
	var stuName string
	fmt.Print("请输入新名字：")
	fmt.Scan(&stuName)
	//保存
	stuObj.name = stuName
	s.allStudent[stuID] = stuObj

	fmt.Println("修改成功！")

}
