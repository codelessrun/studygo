package main

import "fmt"

func main() {
	var n = 100
	fmt.Printf("%T\n", n) // 类型
	fmt.Printf("%v\n", n) // 变量的值，任意类型
	fmt.Printf("%b\n", n) // 二进制
	fmt.Printf("%d\n", n) // 数值
	fmt.Printf("%o\n", n) // 八进制
	fmt.Printf("%x\n", n) // 十六进制

	var s1 = "hello 沙河"
	fmt.Printf("%s\n", s1) // 字符串
	fmt.Printf("%v\n", s1)
	fmt.Printf("%#v\n", s1) // %#v 将值加个描述符，比如字符串自动加上 ”“
}
