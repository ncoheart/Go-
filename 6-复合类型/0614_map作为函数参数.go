package main

import "fmt"

func demo1(m map[int]string) {
	m[4] = "黄月英"
	fmt.Println(m)
}

func main() {
	m := map[int]string{1: "貂蝉", 2: "大乔", 3: "小乔"}
	demo1(m)
	fmt.Println(m)
	//地址传递
}
