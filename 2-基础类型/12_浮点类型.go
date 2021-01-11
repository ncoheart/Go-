package main

import "fmt"

func main() {
	var a float64 = 123.456555555555555
	fmt.Println(a)
	fmt.Printf("%T\n\n", a)

	var b float32 = 123.4565555555555555
	fmt.Println(b) //123.45656
	fmt.Printf("%T\n\n", b)

	c := 123.456
	fmt.Println(c)
	fmt.Printf("%T\n\n", c) //float64
}
