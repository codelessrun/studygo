package main

import "fmt"

const PI = 3.1415926

const (
	n1 = 100
	n2 // 批量声明，如果某一行没有赋值，默认和上面一样
	n3
)

// iota

const (
	a1 = iota
	a2 // a2 = a1 = iota = 1
	a3
)

const (
	b1 = iota
	b2
	_
	b3
)

const (
	c1 = iota // 0
	c2 = 100  // 100
	c3 = iota // 2
	c4        // 3
)

const (
	d1, d2 = iota + 1, iota + 2 // d1 = 1 , d2 = 2
	d3, d4 = iota + 1, iota + 2 // d3 = 2 , d4 = 3
)

// 定义数量级
const (
	_  = iota
	KB = 1 << (10 * iota) // 1左移10位， 2的10次方
	MB
	GB
	TP
	PB
)

func main() {
	//fmt.Println(PI)
	//fmt.Println(n1)
	//fmt.Println(n2)
	//fmt.Println(n3)

	//fmt.Println(a1)
	//fmt.Println(a2)
	//fmt.Println(a3)

	//fmt.Println(b3)

	//fmt.Println(c1)
	//fmt.Println(c2)
	//fmt.Println(c3)
	//fmt.Println(c4)

	//fmt.Println(d1)
	//fmt.Println(d2)
	//fmt.Println(d3)
	//fmt.Println(d4)

	fmt.Println(KB)
	fmt.Println(MB)
	fmt.Println(GB)
	fmt.Println(TP)
	fmt.Println(PB)
}
