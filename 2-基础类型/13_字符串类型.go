package main

import "fmt"

func main() {
	var str1 string //字符串声明
	str1 = "abc"
	fmt.Println(str1)

	str2 := "南京邮电大学"
	fmt.Println(str2) //南京邮电大学\0 \0是字符串结束标志

	//len函数
	fmt.Println(len(str2)) //18，在go语言中每个中文占3个

	//字符串拼接
	str3 := str1 + str2
	fmt.Println(str3) //abc南京邮电大学

}
