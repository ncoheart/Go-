package main

import "fmt"

//递归函数 调用函数自己本身
//递归函数相同的结构
func test01(a int) {
	if a == 1 {
		fmt.Println(a)
		return
	}
	test01(a - 1)
	fmt.Println(a)
}

func test02(num int) int {
	if num == 1 {
		return 1
	}
	return num + test02(num-1)
}

func main() {
	//test01(3)
	fmt.Println(test02(100))
}
