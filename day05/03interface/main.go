package main

import "fmt"

// 定义一个能叫的类型
type speaker interface {
	speak() // 只要实现了speak方法的的变量，都是speaker类型
}

type cat struct {
}

type dog struct {
}

func (c cat) speak() {
	fmt.Println("喵喵喵")
}

func (d dog) speak() {
	fmt.Println("汪汪汪")
}

func da(speaker speaker) {
	// 接收一个参数，传过来什么，我就打什么
	speaker.speak()
}

func main() {
	var c1 cat
	var d1 dog

	da(c1)
	da(d1)

	var ss speaker
	ss = c1 // 多态
	ss = d1
	ss.speak()
}
