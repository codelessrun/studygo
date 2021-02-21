package main

import "fmt"

// 结构体是值类型

// go语言中函数传参永远是副本
type person struct {
	name, gender string
}

func f(x person) {
	x.gender = "女" // 修改的是副本的gender
}

func f2(x *person) {
	(*x).gender = "女"
}

func main() {
	var p person
	p.name = "周林"
	p.gender = "男"

	f(p)
	fmt.Println(p.gender)
	f2(&p)
	fmt.Println(p.gender)

	// 结构体指针2
	var p2 = new(person)
	(*p2).name = "理想"
	p2.gender = "男"

	fmt.Printf("%T\n", p2)
	fmt.Printf("%p\n", p2)  // p2保存的值就是一个内存地址
	fmt.Printf("%p\n", &p2) // p2自己的内存地址

	//	结构体指针3
	// 2.2 使用key-value初始化
	var p3 = &person{
		name:   "元帅",
		gender: "男",
	}
	fmt.Printf("%#v\n", p3)

	// 2.3 使用键值列表形式初始化，值的顺序必须要和结构体的定义顺序一致
	p4 := &person{
		"小王子",
		"男",
	}
	fmt.Printf("%#v\n", p4)

}
