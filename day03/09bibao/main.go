package main

import "fmt"

// 闭包的应用

func f1(f func()) {
	fmt.Println("this is f1")
	f()
}

func f2(x, y int) {
	fmt.Println("this is f2")
	fmt.Println(x + y)
}

// 要求: 使用f1调用f2
// f1(f2)

func f3(f func(int, int), x, y int) func() {
	return func() {
		// 调用f2
		f(x, y)
	}
}

func main() {
	ret := f3(f2, 100, 200)
	f1(ret)
}
