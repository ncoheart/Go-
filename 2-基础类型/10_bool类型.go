package main

import "fmt"

func main() {
	//1.声明变量，默认为false
	var a bool
	fmt.Println(a)

	a = true
	fmt.Println(a)

	//2.布尔类型不接受其他类型的赋值，不支持自动或强制的类型转换
	// a = 1

	//3.自动推导类型
	var b = true
	fmt.Println(b)

	//4.
	c := true
	fmt.Println(c)

	var i, j = 1, 2
	j--
	v2 := i == j
	fmt.Println(v2)
}
