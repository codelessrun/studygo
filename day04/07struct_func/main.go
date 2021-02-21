package main

import "fmt"

// 构造函数

type person struct {
	name string
	age  int
}

// 构造函数
// 返回的是结构体还是结构体指针
// 当结构体比较大的时候，尽量使用结构体的指针，减少程序的内存开销
func newPerson(name string, age int) *person {
	return &person{
		name: name,
		age:  age,
	}
}

func main() {
	p1 := newPerson("元帅", 18)
	p2 := newPerson("周林", 9000)
	fmt.Println(p1, p2)
}
