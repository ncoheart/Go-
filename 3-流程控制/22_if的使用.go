package main

import "fmt"

func main() {
	//变量的初始化
	//var score int = 700
	//自动推导类型
	score := 700

	//if语法结构
	if score == 700 {
		fmt.Println("去上清华北大")
	}

	//if支持一个初始化语句，初始化语句和判断条件用分号分割
	if a := 700; a == 700 {
		fmt.Println("这也是支持的")
	}
}
