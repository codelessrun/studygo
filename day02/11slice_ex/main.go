package main

import (
	"fmt"
	"sort"
)

// 切片的练习
func main() {
	var a = make([]int, 5, 10) // 创建切片，长度为5，容量为10
	fmt.Println(a)
	for i := 0; i < 10; i++ {
		a = append(a, i)
	}
	// 追加是往后追加
	fmt.Println(a)
	fmt.Println(cap(a))

	a1 := [...]int{3, 7, 8, 9, 1}
	sort.Ints(a1[:])
	fmt.Println(a1)
}
