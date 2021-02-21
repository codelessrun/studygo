package main

import "fmt"

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func main() {
	a := 1
	b := 2

	defer calc("1", a, calc("10", a, b))
	a = 0
	defer calc("2", a, calc("20", a, b))
	b = 1

	// 猜结果??

	//	1. a:=1
	//	2. b:=2
	//	3. defer calc("1", 1, calc("10", 1, 2))
	//	4. calc("10", 1, 2) // 10 1 2 3
	//	5. defer calc("1", 1, 3)
	//	6. a=0
	//	7. defer calc("2", 0, calc("20", 0, 2))
	//	8. calc("20", 0, 1) // 20 0 2 2
	//	9. defer calc("2", 0, 2)
	//	10. b=1
	//	11. calc("2", 0, 2) // 2 0 2 2
	//	12. calc("1", 1, 3) // 1 1 3 4

	//  defer 在注册的时候，会把值一起压到内存里。只不过最后再调用，最后调用的时候就不用再取寻值。所以如果defer的参数中有函数，会先把函数算出来，再注册defer
}
