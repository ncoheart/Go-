package main

import "fmt"

type student0 struct {
	id    int
	name  string
	score float64
}

func main() {
	//结构体数组
	var arr [3]student0 = [3]student0{
		student0{
			id:    1,
			name:  "libai",
			score: 98.7,
		},
		student0{
			id:    2,
			name:  "dufu",
			score: 100,
		},
		student0{
			id:    3,
			name:  "baijuyi",
			score: 59.5,
		},
	}
	//fmt.Println(arr)

	//循环打印结构体信息
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	arr[2].score = 99.5
	fmt.Println(arr)

	//切片
	arr0 := []student0{{1, "libai", 89}}
	arr0 = append(arr0, student0{4, "yuanhua", 89})
	fmt.Println(arr0)
}
