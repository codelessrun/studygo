package main

import "fmt"

// 为切片追加元素

func main() {
	s1 := []string{"北京", "上海", "深圳"}
	//s1[3] = "广州" // 索引超出
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))

	// 调用 append 函数必须用原来的切片变量接收返回值
	// append 函数追加元素的时候，如果底层数组放不下的时候，Go底层就会把底层数组换一个，
	// 必须要用原来的变量去接收新的数组位置
	// 扩容策略需要去看一下，不是单纯的翻倍
	s1 = append(s1, "广州")
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))
	s1 = append(s1, "杭州", "成都")
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))

	ss := []string{"武汉", "西安", "苏州"}
	s1 = append(s1, ss...) // ... 表示拆开，因为append只能接收同一元素类型
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))

}
