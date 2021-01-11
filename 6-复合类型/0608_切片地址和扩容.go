package main

import "fmt"

func main() {
	//切片名本身就是一个地址
	//空切片指向内存地址编号为0的空间
	var slice []int
	fmt.Printf("%p\n", slice)

	slice = append(slice, 1, 2, 3)
	fmt.Printf("%p\n", slice)
	fmt.Printf("%p\n", &slice[0])

	slice = append(slice, 4, 5, 6)
	fmt.Printf("%p\n", slice)

	//扩容
	s := make([]int, 5, 8)
	s[0] = 1
	fmt.Println(s)                                    //[1 0 0 0 0]
	fmt.Printf("len = %d,cap = %d\n", len(s), cap(s)) //len = 5,cap = 8
	s = append(s, 1)
	s = append(s, 2)
	s = append(s, 3)
	fmt.Println(s)                                    //[1 0 0 0 0 1 2 3]
	fmt.Printf("len = %d,cap = %d\n", len(s), cap(s)) //len = 8,cap = 8
	s = append(s, 4)
	fmt.Println(s)                                    //[1 0 0 0 0 1 2 3 4]
	fmt.Printf("len = %d,cap = %d\n", len(s), cap(s)) //len = 9,cap = 16

	s1 := make([]int, 0, 1)
	oldcap := cap(s1)
	for i := 0; i < 20000; i++ {
		s1 = append(s1, i)
		newcap := cap(s1)
		if oldcap < cap(s1) {
			fmt.Printf("cap:%d === %d\n", oldcap, newcap)
			oldcap = newcap
		}
	}
	//容量小于1024时，每次扩为原来的两倍
	//容量大于1024时，每次扩容的量为上次的1/4
	/*
		...
		cap:256 === 512
		cap:512 === 1024
		cap:1024 === 1280
		cap:1280 === 1696
		cap:1696 === 2304
		cap:2304 === 3072
		...
	*/

}
