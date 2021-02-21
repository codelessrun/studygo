package main

import "fmt"

// Go 语言中的 return 不是原子操作，在底层是分为2步执行
// 第一步： 函数赋值
// defer 语句
// 第二步： 真正的 RET 返回
// 函数中如果存在 defer ,那么 defer 执行的时机是在第一步和第二部之间

// 没有命名的返回值
func f1() int {
	x := 5 // x 已经赋值了
	defer func() {
		x++ // 修改的的x不是返回值
	}()
	return x
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5 //  返回值 = x
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x // 返回值 =y =x =5
}

func f4() (x int) {
	defer func(x int) {
		x++ // 改变的是x的副本
	}(x)     // 传参数
	return 5 // 返回值 =x = 5
}

func f5() (x int) {
	defer func(x int) int {
		x++
		return x
	}(x)
	return 5
}

func f6() (x int) {
	defer func(x *int) {
		*x++
	}(&x)
	return 5 // 1. 返回值=x=5 ,2 defer x =6 ,3 RET 返回
}

func main() {
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())
	fmt.Println(f5())
	fmt.Println(f6())
}
