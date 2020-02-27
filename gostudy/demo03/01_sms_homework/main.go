package main

import (
	"fmt"
	"os"
)

// 学生管理系统

type student struct {
	id   int64
	name string
}

// StudentMgr 学生管理系统
type StudentMgr struct {
	allStudent map[int64]*student
}

func newStudent(id int64, name string) *student {
	return &student{
		id, name,
	}
}

var smr = StudentMgr{allStudent: make(map[int64]*student, 50)}

func main() {
	printUsage()
}

func printUsage() {
	for {
		fmt.Println(`欢迎访问学生管理系统
		1. 查看所有学生
		2. 新增学生
		3. 删除学生
		4. 退出
		请输入你的选项：`)
		var choice int
		fmt.Scanln(&choice)
		fmt.Printf("你选择了%d个选项！\n", choice)

		switch choice {
		case 1:
			smr.showStudent()
		case 2:
			smr.addStudent()
		case 3:
			smr.deleteStudent()
		case 4:
			os.Exit(1)
		default:
			fmt.Println("out!")
		}
	}
}

func (s *StudentMgr) showStudent() {
	for _, v := range s.allStudent {
		fmt.Printf("学号： %d，姓名： %s\n", v.id, v.name)
	}
}

func (s *StudentMgr) addStudent() {
	var (
		id   int64
		name string
	)
	fmt.Print("请输入学生的学号：")
	fmt.Scanln(&id)
	fmt.Print("请输入学生的姓名：")
	fmt.Scanln(&name)
	newStu := newStudent(id, name)
	s.allStudent[id] = newStu
}

func (s *StudentMgr) deleteStudent() {
	var id int64
	fmt.Print("请输入学生的学号：")
	fmt.Scanln(&id)
	delete(s.allStudent, id)
}
