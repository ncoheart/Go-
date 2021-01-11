package main

import "fmt"

func main() {
	score := 700

	if score >= 680 {
		fmt.Println("可以上清华")
		if score >= 680 {
			fmt.Println("学汽修")
			if score > 700 {
				fmt.Println("学美容美发")
				if score > 710 {
					fmt.Println("学挖掘机")
				}
			}
		}
	} else {
		fmt.Println("不能上清华")
	}
}
