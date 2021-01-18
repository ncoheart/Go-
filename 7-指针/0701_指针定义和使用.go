package main

import "fmt"

func main() {
	var i int
	i = 100
	fmt.Println(i)
	fmt.Printf("%d\n", i)
	fmt.Printf("%p\n", &i)

	var p *int //默认为nil
	p = &i
	fmt.Println(p)

	*p = 1
	fmt.Println(i)

	var p1 *int
	p1 = new(int) //动态分配空间，会自动释放
	*p1 = 57
	fmt.Println(*p1)
}
