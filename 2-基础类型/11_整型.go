package main

import (
	"fmt"
	"unsafe"
)

func main() {
	//uint无符号整型 int有符号整型
	var a = -10
	fmt.Println(a)

	// var b uint = -10 会报错
	// fmt.Println(b)

	b := 20
	fmt.Printf("%T\n", b)         //int
	fmt.Println(unsafe.Sizeof(b)) //8

	var c int32 = 10
	fmt.Printf("%T\n", c)         //int32
	fmt.Println(unsafe.Sizeof(c)) //4
}
