package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Int())    //生成较大的随机数
	fmt.Println(rand.Intn(10)) //不包括10

}
