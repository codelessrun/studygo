package main

import "fmt"

func main() {
	//	for i := 0; i < 10; i++ {
	//		if i == 2 {
	//			goto breakFlag
	//		}
	//		fmt.Println(i)
	//	}
	//breakFlag:
	//	fmt.Printf("结束for循环")

	// 跳出多层循环
	//flag := false
	//for i := 0; i < 10; i++ {
	//	for j := 'A'; j < 'Z'; j++ {
	//		if j == 'C' {
	//			flag = true
	//			break
	//		}
	//		fmt.Printf("%v-%c\n", i, j)
	//	}
	//	if flag {
	//		break
	//	}
	//}

	// goto + label 跳出循环
	for i := 0; i < 10; i++ {
		for j := 'A'; j < 'Z'; j++ {
			if j == 'C' {
				goto xx
			}
			fmt.Printf("%v-%c\n", i, j)
		}
	}

xx: //label 标签
	fmt.Println("over")
}
