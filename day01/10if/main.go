package main

import "fmt"

func main() {
	//age := 19

	//if age > 18 {
	//	fmt.Println("111")
	//} else {
	//	fmt.Println("222")
	//}

	//if age > 35 {
	//	fmt.Println("人到中年")
	//} else if age > 18 {
	//	fmt.Println("青年")
	//} else {
	//	fmt.Println("少年")
	//}

	// 特殊写法，这个时候age作用域只在if中
	if age := 19; age > 18 {
		fmt.Println("111")
	} else {
		fmt.Println("222")
	}
}
