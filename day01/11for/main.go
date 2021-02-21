package main

import "fmt"

func main() {
	// 基本格式
	//for i := 0; i < 10; i++ {
	//	fmt.Println(i)
	//}

	// 变种1,初始语句可省略
	//var i = 5
	//for ; i < 10; i++ {
	//	fmt.Println(i)
	//}

	// 变种2，结束语句也可以省略
	//var i = 5
	//for i < 10 {
	//	fmt.Println(i)
	//	i++
	//}

	// 无限循环
	//for {
	//	fmt.Println("111")
	//}

	// for range (键值循环)
	//s := "hello沙河"
	//for i, v := range s {
	//	fmt.Printf("key: %d value: %c\n", i, v)
	//}

	jiujiu()
}

// 打印九九乘法表
func jiujiu() {
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d ", i, j, i*j)
		}
		// 换行
		fmt.Println()
	}
}
