package main

import "fmt"

// 指针
// GO里面不存在对指针操作，只支持以下两中
// & 取变量地址
// * 根据地址取值
func main() {
	//n := 18
	//fmt.Println(&n)
	//p := &n
	//fmt.Printf("%T %v\n", p, p) // 类型是*int ，代表int类型的指针。

	// 2、*: 根据地址取值
	//m := *p
	//fmt.Println(m)

	//
	//var a *int  // 没有赋值，指针类型的零值是nil
	//*a = 100
	//fmt.Println(a)

	var a1 *int
	fmt.Println(a1)
	var a2 = new(int) // new函数申请一个内存地址，并把地址存在a2中，此时a2应该是个int类型的指针
	fmt.Println(a2)
	fmt.Println(*a2)
	fmt.Printf("%T\n", a2)
	*a2 = 100 // 将a2所指向的地址的值，修改为100
	fmt.Println(*a2)
	fmt.Println(a2)


}
