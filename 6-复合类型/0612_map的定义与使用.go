package main

import "fmt"

func main() {
	name := []string{"张三", "李四", "王五"}
	fmt.Println(name[2])

	//键 & 值
	//map[键类型]值类型{}
	var m map[int]string
	fmt.Println(m)

	//键是唯一的
	//map是无序的列表
	m2 := make(map[int]string, 3) //容量可省略
	m2[1] = "张三"
	m2[2] = "张s"
	m2[3] = "张三"
	m2[4] = "张4"
	//容量是一个个扩容的 len(m2)

	fmt.Println(m2) //map[1:张三 2:张s 3:张三 4:张4]

	//初始化
	m3 := map[int]string{1: "张三", 3: "李四", 9: "王五"}
	fmt.Println(m3[3]) //李四

	//内存编号为0的空间，系统占用不允许读写
}
