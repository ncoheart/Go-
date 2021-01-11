package main

import "fmt"

func main() {
	var a = 10
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	//	int

	//:=也是自动推导，此变量不能被声明过
	b := 10
	b = 20
	fmt.Println(b)

	c := 3.14
	fmt.Printf("%T\n", c)
	//	float64

	d, f, e := 20, 3.14, 30
	fmt.Println(d, f, e)
	fmt.Printf("d:%T f:%T e:%T\n", d, f, e)
}
