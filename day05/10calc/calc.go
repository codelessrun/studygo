package calc

import "fmt"

func init() {
	fmt.Println("import我时自动调用！")
}

// 包中的标识符: 变量名\函数名\结构体\接口等，如果首字母是小写的，表示私有
// 首字母大写的可以被外部包访问
func Add(x, y int) int {
	return x + y
}
