package main

import "fmt"

// 变量的作用域

var x = 100

func f1() {
	// x := 10
	// 函数中查找变量的顺序
	// 1. 现在函数内部查找
	// 2. 找不到就往函数的外面查找，一直找到全局
	fmt.Println(x)
}

func main() {
	f1()

	// 局部变量
	if i := 10; i < 18 {
		fmt.Println("乖乖上线")
	}

	for j := 0; j < 10; j++ {

	}
}
