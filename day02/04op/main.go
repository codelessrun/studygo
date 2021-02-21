package main

import "fmt"

// 运算符

func main() {

	var (
		a = 5
		b = 2
	)

	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)
	a++ // 单独的语句，不能放在等号的右边赋值
	b--

	// 关系运算符, go语言是强类型语言，相同类型才能比较
	fmt.Println(a == b)
	fmt.Println(a != b)
	fmt.Println(a >= b)
	fmt.Println(a > b)
	fmt.Println(a <= b)
	fmt.Println(a < b)

	// 逻辑运算符
	age := 20
	if age > 18 && age < 60 {
		fmt.Println("苦逼上班...")
	} else {
		fmt.Println("不用上班")
	}

	// 位运算, 针对二进制数
	// 5的二进制 101
	// 2的二进制  10

	// &: 按位与（两位均为1,则为1）
	fmt.Println(5 & 2) // 000
	// |: 按位或 (两为有一个1,则为1)
	fmt.Println(5 | 2) // 111
	// ^: 异或 (两为不一样，则为1)
	fmt.Println(5 ^ 2) // 111
	// <<: 左移,讲二进制数左移指定位数 (左移是要在末尾补0)
	fmt.Println(5 << 1)  // 101 ==> 1010
	fmt.Println(1 << 10) // 1 ==> 1000000000 // 1024
	// <<: 右移,将二进制数右移指定位数 (右移是要在末尾去掉多余的)
	fmt.Println(5 >> 1) // 101 ==> 10

	var m = int8(1)
	fmt.Println(m << 10) // 超出，只能存8位

	// 赋值运算
	var x int
	x = 10
	x += 1 // x = x + 1
	x -= 1 // x = x - 1
	x *= 1
	x /= 2 // x= x / 2

	x <<= 2 // x = x << 2
	// ...

}
