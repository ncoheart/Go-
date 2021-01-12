package main

import "fmt"

func main() {
	s := []int{10, 20, 30, 40, 50}

	//截取
	slice := s[0:3:4]       //开始为0，结束为3-1，容量为4(容量不能超过原来切片的容量，可省略)
	fmt.Println(slice)      //[10 20 30]
	fmt.Println(cap(slice)) //5

	//slice = s 和 slice = s[:] 相同
	slice = s[:4]
	fmt.Println(slice) //[10 20 30 40]
	slice = s[2:]
	fmt.Println(slice) //[30 40 50]

	//切片截取后地址还是指向原来切片的部分地址
	//原始切片修改后会影响截取后的切片
}
