package main

import "fmt"

// 方法
// 标识符: 变量名 函数名 类型名 方法名
// Go语言中，如果标识符首字母是大写的，就表示对外部包可见（暴露的，公有的）

type dog struct {
	name string
}

type person struct {
	name string
	age  int
}

func newDog(name string) dog {
	return dog{
		name: name,
	}
}

func newPerson(name string, age int) *person {
	return &person{
		name: name,
		age:  age,
	}
}

// 方法是作用于特定类型的函数
// 接收者表示的是调用该方法的具体类型变量，多用类型名首字母小写表示
func (d dog) wang() {
	fmt.Printf("%s汪汪汪~", d.name)
}

// 注意指针接收者和值接收者的区别
// 什么时候应该使用指针接收者？
// 1. 需要修改接收者中的值
// 2. 接收者是拷贝比较大的对象
// 3. 保证一致性，如果某个方法使用了指针接收者，那么其他的方法也应该使用指针接收者
func (p *person) guonian() {
	p.age++
}

func (p *person) dream() {
	fmt.Println("不上班也有钱拿")
}

func main() {
	d1 := newDog("zhoulin")
	d1.wang()

	p1 := newPerson("元帅", 18)
	fmt.Println(p1.age)
	p1.guonian()
	fmt.Println(p1.age)
	p1.dream()
}
