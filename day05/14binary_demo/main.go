package main

import "fmt"

// 二进制的实际用途

const (
	eat   int = 4 // 100
	sleep int = 2 // 010
	da    int = 1 // 001
)

// 111
// 左边的1表示吃饭     100
// 中边的1表示睡觉     010
// 右边的1表示打豆豆    001

func f(arg int) {
	fmt.Printf("%b\n", arg)
}

// 判断一个数v，是否可以吃饭睡觉打豆豆

// v 要判断的数
// 期望,比如吃饭，睡觉，打豆豆。
// 如果是否符合期望
func f2(v int, expect int) (bol bool) {
	// 如果要判断一个数(v),某一位上是否有1, 那么只要设置这个数(expect),将这个数的相对应的位置设为1
	// 然后只要与这个数(expect)相与，如果结果等于这个数(expect)，那么说明符合条件
	//bol = (v & expect) == expect

	// 等价于
	// 因为任何数和0与都等于0,和1相与，如果是1则为1，如果是0则为0。
	// 那么只需要设置一个数(expect),将这个数的特定位置设为1,其他位置设为0
	// 然后将要比较的值(v)和这个数(expect)相与。
	// 那么除非被比较值(v)对应位置上有1,否则结果将为0
	// 所以只需要判断是否等于0即可判断是否和预期(expect)相等
	bol = (v & expect) != 0
	return
}

func main() {
	//f(eat | da) // 吃饭并且打豆豆
	//f(eat | sleep | da)

	fmt.Println(f2(eat, sleep))
	fmt.Println(f2(eat|da, eat))
	fmt.Println(f2(eat|sleep, da))
}
