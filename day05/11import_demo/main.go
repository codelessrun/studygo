package main

import (
	"fmt"

	zhouolin "code.oldboyedu.com/studygo/day05/10calc"
)

var x = 100

const pi = 3.14

func init() {
	fmt.Println("自动执行")
	fmt.Println(x, pi)
}

func main() {
	ret := zhouolin.Add(10, 20)
	fmt.Println(ret)
}
