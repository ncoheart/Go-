package main

import "fmt"

func main() {
	a, b, c := 54, 30, 20

	sum := a + b + c
	//类型转换 用来在不同但相互兼容的类型之间
	//1.数据类型（变量） int(a)
	//2.数据类型（表达式）
	fmt.Println(float64(sum) / 3)

	var flag bool
	flag = true
	//fmt.Println(int(flag))
	//flag = bool(1) 都会报错
	fmt.Println(flag)

	//在类型转换时建议低类型转为高类型 保证数据的精度
	//int8---int16---int32---int64
	//float32---float64
	//int64----float64
	var d float32 = 3.1
	var f float64 = 3.5
	num := float64(d) + f
	fmt.Println(num)

	//高类型---低类型，丢失精度
	var g int = 1234
	fmt.Println(int8(g)) //-46
}
