package main

import "fmt"

func main() {
	// 跳出 break
	//for i := 0; i < 10; i++ {
	//	if i == 5 {
	//		break
	//	}
	//	fmt.Println(i)
	//}
	//fmt.Println("over")

	// 跳过，continue
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Printf("%d ", i)
	}

}
