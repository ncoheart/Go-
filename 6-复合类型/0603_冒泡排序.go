package main

import "fmt"

func main() {
	arr := [10]int{9, 1, 4, 2, 3, 6, 5, 7, 8, 10}

	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	fmt.Println(arr)
}
