package main

import "fmt"

func main() {
	var a1 = 100
	fmt.Printf("%d\n", a1)
	fmt.Printf("%o\n", a1) // 以八进制输出
	fmt.Printf("%x\n", a1) // 以十六进制输出
	fmt.Printf("%b\n", a1) // 以二进制输出

	// 八进制
	i2 := 077
	fmt.Printf("%d\n", i2)

	// 定义十六进制
	i3 := 0x1234567
	fmt.Printf("%d\n", i3)
	fmt.Printf("%T\n", i3) // 查看变量的类型

	// 明确指定类型
	i4 := int8(9)
	fmt.Printf("%T\n", i4)

}
