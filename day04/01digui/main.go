package main

import "fmt"

// 递归一定要有一个明确的推出条件
// 递归适用于处理那种问题相同、问题的规模越来越小的场景

// 计算阶乘
func f(n uint64) uint64 {
	if n <= 1 {
		return 1
	}
	return n * f(n-1)
}

// 上台阶面试题，一次可以走一步，一次也可以走两步，有多少种走法
func taijie(n uint64) uint64 {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return taijie(n-1) + taijie(n-2)
}

func main() {
	ret := f(5)
	fmt.Println(ret)

	ret2 := taijie(3)
	fmt.Println(ret2)
}
