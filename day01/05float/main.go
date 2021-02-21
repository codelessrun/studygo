package main

import "fmt"

func main() {
	//math.MaxFloat32 float32的最大值

	f1 := 1.23456
	fmt.Printf("%T\n", f1) //默认go语言中的小数都是float64位的
	f2 := float32(1.23456)
	fmt.Printf("%T\n", f2)
	//f1 = f2 // 不可以直接赋值，因为类型不同

}
