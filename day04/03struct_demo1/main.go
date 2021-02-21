package main

import "fmt"

// 结构体

type person struct {
	name   string
	age    int
	gender string
	hobby  []string
}

func main() {
	var p person
	p.name = "周林"
	p.age = 9000
	p.gender = "男"
	p.hobby = []string{"篮球", "足球", "双色球"}
	fmt.Println(p)
	fmt.Printf("%T\n", p)

	// 匿名结构体，多用于临时创建
	var s struct {
		x string
		y int
	}
	s.x = "嘿嘿嘿"
	s.y = 100
	fmt.Printf("%T %v\n", s, s)
}
