package main

import "fmt"

//结构体定义在函数外
//定义
/*type 结构体名字 struct {
	//结构体成员列表
	//成员名 数据类型
	name string
	age int
}*/
type student struct {
	id   int
	name string
	sex  string
	age  int
	addr string
}

func main() {
	//顺序初始化，每个成员必须初始化
	var s1 student = student{1, "xiaoming", "male", 18, "city"}
	//自动推导，指定成员赋值
	s2 := student{
		id:   2,
		name: "xiaohong",
		sex:  "female",
		age:  18,
		addr: "city",
	}
	fmt.Println(s1, s2)

	//定义结构体变量
	var s3 student
	s3.id = 1
	s3.name = "xiaoming"
	s3.age = 18
	s3.sex = "male"
	s3.addr = "city"
	fmt.Println(s3)
	fmt.Printf("%p\n", &s3)
	fmt.Printf("%p\n", &s3.id) //两者地址相同

	//结构体比较
	fmt.Println(s1 == s2) //false
	fmt.Println(s1 == s3) //true 内容一样即为相同
}
