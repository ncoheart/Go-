package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	ID   int //'json:"id"' //通过
	Name string
}

type Class struct {
	Title   string
	Student []Student
}

func newStudent(id int, name string) Student {
	return Student{
		ID:   id,
		Name: name,
	}
}

func main() {
	// 结构体字段大写开头表示可公开访问
	// 小写表示私有（仅在定义当前结构体的包中可以访问）

	c1 := Class{
		Title:   "火箭101",
		Student: make([]Student, 0, 20),
	}

	for i := 0; i < 10; i++ {
		tmpStu := newStudent(i, fmt.Sprintf("stu%02d", i))
		c1.Student = append(c1.Student, tmpStu)
	}
	// JSON序列化
	// 如果class结构体中title首字母没有大写，json包是访问不了的
	data, err := json.Marshal(c1)
	if err != nil {
		fmt.Println("json marshal failed, error:", err)
	}
	fmt.Printf("%s\n", data)

	// 反序列化
	jsonStr := "{\"Title\":\"火箭101\",\"Student\":[{\"ID\":0,\"Name\":\"stu00\"},{\"ID\":1,\"Name\":\"stu01\"},{\"ID\":2,\"Name\":\"stu02\"},{\"ID\":3,\"Name\":\"stu03\"},{\"ID\":4,\"Name\":\"stu04\"},{\"ID\":5,\"Name\":\"stu05\"},{\"ID\":6,\"Name\":\"stu06\"},{\"ID\":7,\"Name\":\"stu07\"},{\"ID\":8,\"Name\":\"stu08\"},{\"ID\":9,\"Name\":\"stu09\"}]}"
	var c2 Class
	err = json.Unmarshal([]byte(jsonStr), &c2)
	if err != nil {
		fmt.Println("json unmarshal failed, error:", err)
	}
	fmt.Println(c2)

	//fmt.Printf("%#vn",c1)
}
