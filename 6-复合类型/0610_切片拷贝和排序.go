package main

import "fmt"

func main() {
	var s []int = []int{1, 2, 3, 4, 5}

	s1 := make([]int, 5)
	copy(s1, s)
	fmt.Println(s1)

	//拷贝的切片相互独立

	arr := []int{9, 1, 4, 2, 3, 6, 5, 7, 8, 10}

	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	fmt.Println(arr)
}
