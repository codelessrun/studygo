package main

import "fmt"

// 切片 slice
func main() {
	var s1 []int    // 定义一个存放int的切片
	var s2 []string // 定义一个存放string的切片
	fmt.Println(s1, s2)
	fmt.Println(s1 == nil)
	fmt.Println(s2 == nil)
	// 初始化
	s1 = []int{1, 2, 3}
	s2 = []string{"沙河", "张江", "平台村"}
	fmt.Println(s1, s2)
	fmt.Println(s1 == nil)
	fmt.Println(s2 == nil)

	// 切片的长度和容量
	fmt.Printf("len(s1): %d cap(s1):%d\n", len(s1), cap(s1))
	fmt.Printf("len(s2): %d cap(s2):%d\n", len(s2), cap(s2))

	// 2. 由数组得到切片
	a1 := [...]int{1, 3, 5, 7, 9, 11, 13}
	s3 := a1[0:4] //基于数组切割，左包含右不包含。 // 按索引切割
	fmt.Println(s3)
	s4 := a1[1:6]
	fmt.Println(s4)
	s5 := a1[:4] // ==> [0:4] => [7 9 11 13]
	s6 := a1[3:] // ==> [3:len(a1)]
	s7 := a1[:]  // ==> [0:len(a1)]
	fmt.Println(s5, s6, s7)
	// ** 切片的容量是指底层数组的容量 (不准确)
	fmt.Printf("len(s5): %d cap(s5):%d\n", len(s5), cap(s5))
	// ??? 应该是从切片的底层数组所切的第一个元素到底层数组的数量。切片的底层数组从第一个元素到最后
	fmt.Printf("len(s6): %d cap(s6):%d\n", len(s6), cap(s6))

	// 切片再切片
	s8 := s6[3:]
	fmt.Printf("len(s8): %d cap(s8):%d\n", len(s8), cap(s8))
	fmt.Println("s6: ", s6)
	// 切片是一个引用类型，都指向了一个底层数组
	a1[6] = 1300
	fmt.Println("s6: ", s6)
	fmt.Println("s8: ", s8)
}
