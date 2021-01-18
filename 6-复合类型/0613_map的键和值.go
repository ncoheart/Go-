package main

import "fmt"

func main() {
	//map中key类型必须支持 == != 一般建议写 基本类型
	//切片 函数 切片的结构类型 不能作为字典的key
	//m:=make(map[[]string]int) //error invalid map key type []string
	m := make(map[string][3]int)
	fmt.Println(m)

	m["xiaoming"] = [3]int{100, 99, 98}
	m["xiaoliang"] = [3]int{6, 8, 9}
	m["xiaohong"] = [3]int{100, 100, 100}

	for k, v := range m {
		fmt.Println(k, v)
	}

	// 判断键值对是否存在
	value, ok := m["xiaolei"]
	fmt.Println(value, ok) //[0 0 0] false

	//删除map中的元素，根据key进行删除
	delete(m, "xiaoming")
	fmt.Println(m)

}
