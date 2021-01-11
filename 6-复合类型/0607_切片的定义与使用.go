package main

import "fmt"

func main() {
	//数组的长度是固定的
	arr := [5]int{1, 2, 3, 4, 5}
	//arr[5] = 10
	fmt.Println(arr)

	//切片定义
	var slice []int //空切片
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
	slice = append(slice, 1, 2, 3)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	//2.
	var s1 []int = []int{1, 2, 3}
	fmt.Println(s1)

	//自动推导
	s2 := []int{1, 2, 3, 4, 5}
	fmt.Println(s2)

	//3.
	s := make([]int, 5, 10)   //长度是5，容量是10
	s = append(s, 1, 2, 3, 4) //[0 0 0 0 0 1 2 3 4]
	fmt.Println(s)

	/*可以下标赋值和循环赋值，不能使用容量作为循环条件*/
}
