package main

import "fmt"

func demo(s []int) {
	fmt.Println(s)
	//切片形参和实参是同一个地址，修改会影响原实参的值
	s[0] = 123
}
func main() {
	s := []int{1, 2, 3, 4, 5}
	// 地址传递
	demo(s)
	fmt.Println(s)
	// 切片在内存中存储，如果该空间没有指向，就会被销毁
}
