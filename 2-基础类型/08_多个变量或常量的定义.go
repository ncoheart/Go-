package main

import "fmt"

func main() {
	//	变量：程序运行期间可以改变的两 声明关键字var
	/*	var a int
		a = 10
		fmt.Println(a)

		b,c := 10,20
		fmt.Println(b, c)*/

	//	不同类型变量的定义
	var a = 1
	var b = 3.14
	var (
		c = 3
		d = 2.0
	)
	e, f := 1, 4.56
	fmt.Println(a, b, c, d, e, f)

	//	常量：程序运行期间不可以更改的量，关键字const
	/*	const i int = 10
		const j float64 = 3.14*/

	const (
		m int     = 10
		n float64 = 10.6
	)
	fmt.Println(m, n)
	//	自动推导
	const (
		i = 10
		j = 20
	)
}
