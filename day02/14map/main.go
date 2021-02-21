package main

import "fmt"

func main() {
	// key 是string value 是 int
	var m1 map[string]int
	m1 = make(map[string]int, 10) // 最好先估算容量，避免动态扩容
	m1["理想"] = 18
	m1["姬无命"] = 35
	//fmt.Println(m1)
	//fmt.Println(m1["理想"])

	//fmt.Println(m1["娜扎"]) // 如果不存在这个key的值，则返回对应类型的零值
	// 值和布尔值
	// 约定成俗用ok
	//value, ok := m1["娜扎"]
	//if !ok {
	//	fmt.Println("查无此人")
	//} else {
	//	fmt.Println(value)
	//}

	// map 的遍历
	//for k, v := range m1 {
	//	fmt.Println(k, v)
	//}
	//// 只遍历 key
	//for k := range m1 {
	//	fmt.Println(k)
	//}
	//// 只遍历v
	//for _, v := range m1 {
	//	fmt.Println(v)
	//}

	// 删除
	delete(m1, "姬无命")
	fmt.Println(m1)
	delete(m1, "娜扎") // 删除不存在的
	fmt.Println(m1)

}
