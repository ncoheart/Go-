package main

import "fmt"

func main() {
	var v1 complex64
	v1 = 3.2 + 12i
	v2 := 3.2 + 12i
	v3 := complex(3.2, 12)

	fmt.Println(v1, v2, v3)
	fmt.Println(real(v1), imag(v1))
}
