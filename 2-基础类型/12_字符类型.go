package main

import "fmt"

func main() {
	//1.单引号表示字符，双引号表示字符串
	var a byte = 'a'
	fmt.Println(a)         //97
	fmt.Printf("%c\n", 98) //b

	//\n \t
	fmt.Println('\n') //10
}
