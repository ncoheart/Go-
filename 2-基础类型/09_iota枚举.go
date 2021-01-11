package main

import "fmt"

func main() {
	//1.iota常量生成器，每个一行，自动加1
	//2.iota给常量赋值使用
	const (
		x = iota
		y = iota
		z = iota
	)
	//x=0,y=1,z=3
	fmt.Println(x, y, z)

	//3.iota遇到const，重置为0
	const d = iota
	//d=0
	const e = iota
	fmt.Println(d, e)

	//4.某些情况下，可以省略后面的iota
	const (
		a1 = iota
		b1
		c1
	)
	fmt.Println(a1, b1, c1)

	//5.如果在同一行，则值相同
	const (
		i1         = iota             //0
		j1, j2, j3 = iota, iota, iota //1,1,1
		k1         = iota             //2
	)
	fmt.Println(i1, j1, j2, j3, k1)
}
