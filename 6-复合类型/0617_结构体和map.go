package main

import "fmt"

type hero struct {
	name  string
	age   int
	power int
}

func main() {
	//将结构体作为map中的值 value
	m := make(map[int]hero)
	m[1] = hero{
		name:  "钢铁侠",
		age:   30,
		power: 100,
	}
	m[2] = hero{
		name:  "美国队长",
		age:   320,
		power: 90,
	}
	// map不建议进行排序
	fmt.Println(m)
	delete(m, 2)

	fmt.Println(m)
}
