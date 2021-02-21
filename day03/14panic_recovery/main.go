package main

import "fmt"

func funA() {
	fmt.Println("a")
}

func funB() {
	// 刚刚打开的数据库连接
	defer func() {
		// 修复，让程序继续往下走。
		err := recover() // 尽量少用，改跑错就抛错
		fmt.Println(err)
		fmt.Println("释放数据库连接")
	}()
	panic("出现了严重的错误！")
	fmt.Println("b")
}

func funC() {
	fmt.Println("c")
}

// 注意
// recover() 必须搭配 defer 使用
// defer 一定要在可能引发panic的语句之前定义
func main() {
	funA()
	funB()
	funC()
}
