package main

import "fmt"

func main() {
	//var 数组名 [长度]类型
	//没有初始化的话，默认都为0
	var a [10]int
	fmt.Println(len(a))
	fmt.Println(a) //[0 0 0 0 0 0 0 0 0 0]

	//注意 error
	//var arr int = 10
	//var b[arr]int = 100

	//数组赋值
	for i := 0; i < len(a); i++ {
		a[i] = i
	}
	fmt.Println(a) //[0 1 2 3 4 5 6 7 8 9]
	//fmt.Println(a[10]) //数组索引10 无效(超出 10 元素数组的界限)

	//匿名变量
	for _, v := range a {
		fmt.Println(v)
	}

	//数组初始化
	var b [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println(b)

	//自动推导
	c := [5]int{1, 2, 3, 4, 5}
	fmt.Println(c)

	//部分初始化
	d := [5]int{1, 2, 3}
	fmt.Println(d) //[1 2 3 0 0]

	//指定某个元素初始化
	e := [5]int{2: 10, 4: 20}
	fmt.Println(e) //[0 0 10 0 20]

	//通过初始化确定长度
	f := [...]int{1, 2, 3}
	fmt.Println(len(f)) //3

	/*
		常见的问题：
		1.数组长度
			必须为常量
		2.下标
			注意数组下标越界
			var arr [5]int = [5]int{1,2,3,4,5}
			arr[6]
			arr[-1] //error
		3.数组名
			arr = 123 //error
		4.两个数组类型相同，个数相同 可以复制
			arr1 := arr
			fmt.Printf("%T\n",arr) //[5]int
		5.数组名表示整个数组，数组名对应地址就是数组第一个元素的地址
			fmt.Printf("%p\n", &arr[0]) //地址是连续的
	*/

}
