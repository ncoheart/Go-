package main

import "fmt"

func test11() func() int {
	var x int

	//返回匿名函数
	return func() int {
		x++
		//返回整型
		return x
	}
}
func main() {

	//匿名函数 没有名字的函数
	var num int
	num = 9
	f := func() {
		num++
		fmt.Println(num)
	} //花括号 {}()调用
	//函数调用
	//f()
	fmt.Println(num)

	type FuncType func()

	var F1 FuncType
	F1 = f
	f()
	F1()

	//参数传递
	func(a, b int) {
		var sum int
		sum = a + b
		fmt.Println(sum)
	}(3, 6)

	//匿名函数返回值
	x, y := func(i, j int) (max, min int) {
		if i > j {
			max = i
			min = j
		} else {
			max = j
			min = i
		}
		return
	}(10, 20)

	fmt.Println(x, y)

	//闭包
	f2 := test11() //接收匿名函数
	//闭包在使用时，已经初始化的变量不会重复初始化
	fmt.Println(f2()) //1
	fmt.Println(f2()) //2
	fmt.Println(f2()) //3
}
