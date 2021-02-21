package main

import "fmt"

type animal interface {
	move()
	eat(string)
}

type cat struct {
	name string
	feet int8
}

// 使用值接收者实现了接口的所有方法
//func (c cat) move() {
//	fmt.Println("走猫步...")
//}
//func (c cat) eat(food string) {
//	fmt.Printf("猫吃%s...\n", food)
//}

// 使用指针接收者实现了所有的接口方法
func (c *cat) move() {
	fmt.Println("走猫步...")
}
func (c *cat) eat(food string) {
	fmt.Printf("猫吃%s...\n", food)
}

// 使用值接收者与使用指针接收者实现接口的区别？

func main() {
	var a1 animal

	c1 := cat{"tom", 4}  // cat
	c2 := &cat{"假老练", 4} // *cat

	a1 = &c1 //实现animal这个接口的是这个cat的指针
	fmt.Println(a1)
	a1 = c2
	fmt.Println(a1)
}
