package main

import "fmt"

// make() 函数创造切片
func main() {
	//s1 := make([]int, 5, 10) // 类型，长度，容量
	//s1 := make([]int, 5) // 类型，长度，容量
	//fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))
	//
	//s2 := make([]int, 0, 10)
	//fmt.Printf("s2=%v len(s2)=%d cap(s2)=%d\n", s2, len(s2), cap(s2))

	// 切片不能直接比较，唯一合法的操作是和nil比较。一个nil值的切片并没有底层数组，长度和容量都为0，但是长度和容量都为0的切片不一定是nil

	//var v1 []int  // 声明没有开辟内存空间
	//v2 := []int{} // 以下这两个都开辟了内存空间
	//v3 := make([]int, 0)
	//fmt.Printf("len(v1)=%d cap(v1)=%d v1==nill:%v\n", len(v1), cap(v1), v1 == nil)
	//fmt.Printf("len(v2)=%d cap(v2)=%d v2==nill:%v\n", len(v2), cap(v2), v2 == nil)
	//fmt.Printf("len(v3)=%d cap(v3)=%d v3==nill:%v\n", len(v3), cap(v3), v3 == nil)
	// 所以，要判读一个切片是不是空的，使用len(),不能和nil比较

	// 切片的赋值，因为共享内存，所以改一个，其他的也被改了
	s3 := []int{1, 2, 3}
	s4 := s3
	fmt.Println(s3, s4)
	s3[0] = 1000
	fmt.Println(s3, s4)

	//切片的遍历
	for i := 0; i < len(s3); i++ {
		fmt.Println(s3[i])
	}
	for i, v := range s3 {
		fmt.Println(i, v)
	}

}
