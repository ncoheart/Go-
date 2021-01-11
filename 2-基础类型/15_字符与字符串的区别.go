package main

import "fmt"

func main() {
	var ch byte
	var str string

	ch = 'a'
	str = "a"

	fmt.Println(ch)            //97
	fmt.Println(str)           //a
	fmt.Println(str[0])        //97
	fmt.Printf("%c\n", str[0]) //a
}
