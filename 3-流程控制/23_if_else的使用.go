package main

import "fmt"

func main() {
	var score int = 100
	if score >= 200 {
		fmt.Println("去上大学")
	} else {
		fmt.Println("你想多了")
	}

	if score >= 200 {
		fmt.Println("去上大学")
	} else if score >= 100 {
		fmt.Println("二本也行")
	} else {
		fmt.Println("快跑")
	}
}
