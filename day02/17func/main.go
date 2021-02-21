package main

import "fmt"

func main() {
	//r := sum(1, 2)
	//fmt.Println(r)
	f7("下雨了")
	f7("下雨了", 1, 2, 3)
}

//func sum(x int, y int) (ret int) {
//	return x + y
//}

// 返回值可以命名也可以不命名
// 如果命名了，那么可以直接在函数内使用
func sum(x int, y int) (ret int) {
	ret = x + y
	return
}

// 参数的类型简写,类型一样时，可以省略前面的
func f6(x, y int) int {
	return x + y
}

// 可变长参数
// 可变长参数必须放在函数的最后
func f7(x string, y ...int) {
	fmt.Println(x)
	fmt.Println(y)
}
