package main

import (
	"fmt"
	"os"
)

func showMenu() {
	fmt.Println("欢迎来到学员信息管理系统")
	fmt.Println("1. 添加学员")
	fmt.Println("2. 编辑学员信息")
	fmt.Println("3. 展示所有学员信息")
	fmt.Println("4. 退出系统")
}

func getInput() *student {
	var (
		id    int
		name  string
		class string
	)
	fmt.Println("请按要求输入学员信息")
	fmt.Print("请输入学员的学号：")
	fmt.Scanf("%d\n", &id)
	fmt.Print("请输入学员的姓名：")
	fmt.Scanf("%s\n", &name)
	fmt.Print("请输入学员的班级：")
	fmt.Scanf("%s\n", &class)

	return newStudent(id, name, class)

}

func main() {
	sm := newStudentMgr()
	for {
		showMenu()

		var input int
		fmt.Print("请输入需要操作的序号：")
		fmt.Scanf("%d\n", &input)
		fmt.Println("用户输入的是：", input)

		switch input {
		case 1:
			stu := getInput()
			sm.addStudent(stu)
		case 2:
			stu := getInput()
			sm.modifyStudent(stu)
		case 3:
			sm.showStudent()
		case 4:
			os.Exit(0)
		}
	}

}
