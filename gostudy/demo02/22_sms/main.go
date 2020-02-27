package main

import (
	"fmt"
	"os"
)

// 函数版学生管理系统

type student struct {
	id   int64
	name string
}

func newStudent(id int64, name string) *student {
	return &student{
		id, name,
	}
}

var (
	allStudent map[int64]*student
)

func main() {
	allStudent = make(map[int64]*student, 50)
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
			showStudent()
		case 2:
			addStudent()
		case 3:
			deleteStudent()
		case 4:
			os.Exit(1)
		default:
			fmt.Println("out!")
		}
	}
}

func showStudent() {
	for k, v := range allStudent {
		fmt.Printf("学号： %d，姓名： %s\n", k, v.name)
	}
}

func addStudent() {
	var (
		id   int64
		name string
	)
	fmt.Print("请输入学生的学号：")
	fmt.Scanln(&id)
	fmt.Print("请输入学生的姓名：")
	fmt.Scanln(&name)
	s := newStudent(id, name)
	allStudent[id] = s
}

func deleteStudent() {
	var id int64
	fmt.Print("请输入学生的学号：")
	fmt.Scanln(&id)
	delete(allStudent, id)
}
