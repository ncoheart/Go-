package main

import "fmt"

//函数名在整个项目里唯一
func add(s1, s2 int) {
	sum := s1 + s2
	fmt.Println(sum)
}
func main() {
	a := 10
	b := 20
	//	函数调用
	add(a, b)

	/*
		内存相对于一个可执行程序来说 可以分为四个区
		代码区、数据区、堆区、栈区
		栈区用来存储程序执行过程中函数内部定义的信息和局部变量

	*/
}
