package main

import "fmt"

func main() {
	a1 := [...]int{1, 3, 5, 7, 9, 11, 13, 15, 17}
	s1 := a1[:]

	// 删掉索引为1的那个3
	s1 = append(a1[0:1], a1[2:]...)
	fmt.Println(s1)
	fmt.Println(a1) // ?
}
