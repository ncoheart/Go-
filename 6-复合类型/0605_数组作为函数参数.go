package main

import "fmt"

func BubbleSort(arr [10]int) [10]int {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	fmt.Println(arr)
	return arr
}

func main() {
	/*
		形参和实参是不同地址
		寒素调用结束后，内存回收，不会影响主函数实参的值
	*/
	var arr [10]int = [10]int{9, 1, 4, 2, 3, 6, 5, 7, 8, 10}
	arr = BubbleSort(arr) //[1 2 3 4 5 6 7 8 9 10]
	fmt.Println(arr)      //[1 2 3 4 5 6 7 8 9 10]
}
