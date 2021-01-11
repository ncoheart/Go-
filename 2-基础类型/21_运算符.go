package main

import "fmt"

func main() {
	a := 10
	fmt.Println(a)
	a++ //a自增1
	fmt.Println(a)

	//关系运算符 略
	fmt.Println(a == 11) //true

	x1 := 60 // 0011 1100
	x2 := 13 // 0000 1101
	var y int

	y = x1 & x2 // 0000 1100
	fmt.Println(y)

	fmt.Printf("%p\n", &x1)
	fmt.Printf("%d\n", *&x1) //60

	//运算符的优先级
}
