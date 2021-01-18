package main

import "fmt"

type student struct {
	id    int
	name  string
	class string
}

// 学员管理的类型
type studentMgr struct {
	allStudents []*student
}

func newStudent(id int, name string, class string) *student {
	return &student{
		id:    id,
		name:  name,
		class: class,
	}
}

// studentMgr的构造函数
func newStudentMgr() *studentMgr {
	return &studentMgr{allStudents: make([]*student, 0, 100)}
}

// 1. 添加学生
func (s *studentMgr) addStudent(newStu *student) {
	s.allStudents = append(s.allStudents, newStu)
}

// 2. 编辑学生
func (s *studentMgr) modifyStudent(newStu *student) {
	for i, v := range s.allStudents {
		if newStu.id == v.id {
			s.allStudents[i] = newStu
			return
		}
	}
	fmt.Println("输入的学生信息有误，请检查后重新输入")
}

// 3. 展示学生
func (s *studentMgr) showStudent() {
	for _, v := range s.allStudents {
		fmt.Printf("学号：%d 姓名：%s 班级：%s\n", v.id, v.name, v.class)
	}
}
