package main

import "fmt"

func test() (a, b, c int) {
	return 1, 2, 3
}

func main() {
	/*a := 10
	b := 20
	c := 30*/

	//多重赋值
	a, b, c := 10, 20, 30
	fmt.Println(a, b, c)

	a = b
	b = a
	fmt.Println(a, b)

	i := 10
	j := 20
	i, j = j, i
	fmt.Println(i, j)

	temp, _ := 10, 7
	fmt.Println(temp)

	var _, e, f = test()
	fmt.Println(e, f)
}
