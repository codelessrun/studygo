package main

import "fmt"

//var name string
//var age int
//var isOk bool

var (
	name string
	age  int
	isOk bool
)

func main() {
	//name = "理想"
	//age = 18
	//isOk = true
	//fmt.Printf("name: %s", name)
	//println(age)
	//print(isOk)

	//var s1 string = "王海滨"
	//println(s1)
	//
	//var s2 = "20"
	//println(s2)
	//
	//// 只能在函数里用
	//s3 := "黑黑黑"
	//println(s3)

	x, _ := foo()
	_, y := foo()
	fmt.Println("x=", x)
	fmt.Println("y=", y)
}

func foo() (int, string) {
	return 10, "Q1Mi"
}
