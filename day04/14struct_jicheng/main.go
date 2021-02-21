package main

import "fmt"

// 结构体中模拟实现其他语言中的"继承"

// 动物
type animal struct {
	name string
}

// 给动物加一个移动方法
func (a animal) move() {
	fmt.Printf("%s会动\n", a.name)
}

type dog struct {
	feet uint64
	animal  // animal 拥有的方法，dog此时也有了。
}

// 给狗加一个汪汪汪方法
func (d dog) wang() {
	fmt.Printf("%s在叫:汪汪汪~\n", d.name)
}

func main() {
	d1 := dog{
		feet: 4,
		animal: animal{
			name: "周林",
		},
	}

	d1.wang()
	d1.move()
}
