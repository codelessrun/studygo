package main

import "fmt"

// go 里面数字必须指定容量和长度
// 数组的容量和长度是类型的一部分
func main() {
	// 定义数组
	//var a1 [3]bool
	//var a2 [4]bool
	//fmt.Printf("a1: %T a2: %T\n", a1, a2)
	//
	//// 数组初始化
	//fmt.Println(a1)
	//fmt.Println(a2)
	//// 初始化方式1
	//a1 = [3]bool{true, true, false}
	//// 初始化方式2,
	//// 根据初始值自动推断数组的长度是多少
	//a10 := [...]int{0, 1, 2, 3, 4, 5, 6}
	//fmt.Println(a10)
	//// 初始化方式3,不写默认为零值
	////a3 := [5]int{1, 2}
	//// 根据索引指定
	//a3 := [5]int{0: 1, 4: 2}
	//fmt.Println(a3)

	// 数组的遍历
	//cities := [...]string{"上海", "北京", "广州", "深圳"}
	//for i := 0; i < len(cities); i++ {
	//	fmt.Println(cities[i])
	//}

	// for range 遍历
	//for i, v := range cities {
	//	fmt.Println(i, v)
	//}

	// 多维数组
	// [[1 2],[2 3],[3 4]]
	//var a11 [3][2]int
	//a11 = [3][2]int{
	//	[2]int{1, 2},
	//	[2]int{2, 3},
	//	[2]int{3, 4},
	//}
	//fmt.Println(a11)

	// 多维数组的遍历
	//for _, v1 := range a11 {
	//	fmt.Println(v1)
	//	for _, v2 := range v1 {
	//		fmt.Println(v2)
	//	}
	//}

	// 数组是值类型
	//b1 := [...]int{1, 2, 3}
	//b2 := b1
	//b2[0] = 100
	//fmt.Println(b1)
	//fmt.Println(b2)

	sum := sum()
	fmt.Println(sum)
	f2()
}

// 习题1，求数组的和
func sum() int {
	sum := 0
	array := [...]int{1, 3, 5, 7, 8}
	for v := range array {
		sum = sum + v
	}
	return sum
}

// 习题2, 找出数组中和为指定值的两个元素的下标。比如 [1 3 5 7 8] 中和为8的两个元素的下标分别为（0，3）和（1，2）
func f2() {
	array := [...]int{1, 3, 5, 7, 8}

	//for i := 0; i < len(array); i++ {
	//	for j := 0; j < i; j++ {
	//		if array[i]+array[j] == 8 {
	//			fmt.Println(i, j)
	//		}
	//	}
	//}

	for i := 0; i < len(array); i++ {
		for j := i; j < len(array); j++ {
			if array[i]+array[j] == 8 {
				fmt.Println(i, j)
			}
		}
	}

}
