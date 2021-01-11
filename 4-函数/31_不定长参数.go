package main

import "fmt"

//函数定义 形参
/*func sum(a,b int)  {
	sum := a + b
	fmt.Println(sum)
}*/

//参数名 自己起 ... type 不定参
func sum(args ...int) {
	//sum := args[0]+args[1]+args[2]+args[3]
	//fmt.Println(sum)
	sum := 0
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}
	fmt.Println(sum)
}

//不定参放在列表的最后
func dec(a int, args ...int) {
	fmt.Println(a - args[0])
}
func main() {
	//函数调用
	//sum(1,3) //实参

	sum(1, 2, 3, 4, 4, 3, 21) //[1 2 3 4]
	dec(1, 3, 4, 5, 6, 7)
	//	固定参数必须传值，不定参数根据需要决定
}
