package main

import "fmt"

func main() {
	type bigint int64
	var a bigint
	a = 10
	fmt.Println(a)
	fmt.Printf("%T\n", a) //main.bigint

	type (
		long int64
		char byte
	)

	var b long = 11
	var ch char = 'a'
	fmt.Println(b, ch)
}
