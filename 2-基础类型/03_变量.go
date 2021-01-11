package main

import "fmt"

func main() {
	//	变量的声明
	//	1.声明 var 变量名 类型
	//	2.只是声明变量，默认为0
	//	3.在同一个{}中，声明的变量是唯一的
	var a int
	fmt.Println(a)

	var b, c int
	b, c = 20, 30
	fmt.Println(b + c)

	//	变量的初始化
	var d = 10
	d = 20
	fmt.Println(d)

	var e = 3.14
	fmt.Println(e)
	var value float64 = 2
	//等号后面可以跟表达式
	var sum = value * value
	fmt.Println(sum)
}
