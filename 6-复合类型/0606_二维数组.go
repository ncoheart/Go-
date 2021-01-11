package main

import "fmt"

func main() {
	var arr1 [10]int  //一维数组
	fmt.Println(arr1) //[0 0 0 0 0 0 0 0 0 0]

	var arr2 [2][3]int
	fmt.Println(arr2) //[[0 0 0] [0 0 0]]
	arr2[0][2] = 1
	fmt.Println(arr2) //[[0 0 1] [0 0 0]]
	//行数
	fmt.Println(len(arr2)) //2
	//列数
	fmt.Println(len(arr2[0])) //3

	b := [3][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	fmt.Println(b)

	//部分初始化
	c := [3][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}}
	fmt.Println(c)

}
