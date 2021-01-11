package main

import "fmt"

//go build main.go test.go
func main() {
	fmt.Println("main")
	test() //undefined: test
}
