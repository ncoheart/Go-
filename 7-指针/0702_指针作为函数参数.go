package main

import "fmt"

func Swap(num1, num2 int) {
	num1, num2 = num2, num1
	//fmt.Println(num1,num2)
}

func SwapbyPoint(num1, num2 *int) {
	*num1, *num2 = *num2, *num1
	//fmt.Println(num1,num2)
}

func main() {
	a := 10
	b := 20
	Swap(a, b)
	fmt.Println(a, b)
	SwapbyPoint(&a, &b)
	fmt.Println(a, b)

}
