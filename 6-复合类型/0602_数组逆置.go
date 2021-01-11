package main

import "fmt"

func main() {
	//	练习：从以下一个整数数组中取出最大的整数，最小整数，总和，平均值
	//	var a [6]int = [6]int{0,1,2,3,4,5}
	var a [6]int = [6]int{10, -1, 2, -3, 4, 5}

	var max, min, sum int = a[0], a[0], 0

	for i := 0; i < len(a); i++ {
		if a[i] > max {
			max = a[i]
		}
		if a[i] < min {
			min = a[i]
		}
		sum += a[i]
	}

	fmt.Println(max, min, sum, float64(sum)/float64(len(a)))

	i := 0
	j := len(a) - 1
	for i <= j {
		a[i], a[j] = a[j], a[i]
		i++
		j--
	}
	fmt.Println(a)
}
