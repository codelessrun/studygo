package main

import (
	"fmt"
	"unicode"
)

// byte 和 rune 类型

// GO语言为了处理非ASCII码类型的字符，定义了新的类型 rune

func main() {
	s := "hello 沙河"
	// 求字符的长度
	n := len(s)
	fmt.Println(n)

	//for i := 0; i < len(s); i++ {
	//	//fmt.Println(s[i])
	//	fmt.Printf("%c\n", s[i]) // %c 字符,如何有中文会乱码
	//}

	//for _, c := range s { // 从字符中拿出具体的字符
	//	//fmt.Println(c)c
	//	fmt.Printf("%c\n", c) // 不会乱码
	//}

	// 字符串修改,(字符串不允许直接修改)
	//s2 := "白萝卜"
	//s3 := []rune(s2) // 把字符串强制转换成了一个rune切片
	//s3[0] = '红'
	//fmt.Println(string(s3)) // 把rune切片强制转换成string

	//c1 := "红"
	//c2 := '红' // rune(int32)
	//fmt.Printf("c1: %T c2: %T\n", c1, c2)

	//c3 := "H"
	//c4 := byte('H') // byte(uint8)
	//fmt.Printf("c3: %T c4: %T\n", c3, c4)
	//fmt.Printf("%d\n", c4)

	// 统计汉字数量
	cr := "hello 沙河小王子"
	fmt.Printf("汉字的数量是%v个\n", isHanNum(cr))
}

// 统计字符串中汉字的数量
func isHanNum(str string) int {
	num := 0
	for _, c := range str {
		if unicode.Is(unicode.Han, c) {
			num++
		}
	}
	return num
}
