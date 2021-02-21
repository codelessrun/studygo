package main

import (
	"fmt"
	"strings"
	"unicode"
)

// 统计一个字符串中每个单词出现的次数，比如“how do you do ” how=1,do=2,you=1
func main() {
	str1 := "hello沙河"
	fmt.Println(count(str1))

	str2 := "how do you do"
	fmt.Println(count2(str2))

	str3 := "上海自来水来自海上"
	fmt.Println(IsHuiWen(str3))
}

// 判断一段string中汉字出现的次数
func count(str string) (count int) {
	for _, c := range str {
		if unicode.Is(unicode.Han, c) {
			count++
		}
	}
	return
}

// 统计一段话中单词出现的次数
func count2(str string) map[string]int {
	s1 := strings.Split(str, " ")
	m1 := make(map[string]int, 10)
	for _, w := range s1 {

		// 如果这个变量第一次出现

		//if _, ok := m1[w]; !ok {
		//	m1[w] = 1
		//} else {
		//	// 如果不是第一次出现，则累加
		//	m1[w]++
		//}

		// 因为默认值是0，所有可以直接累加
		m1[w]++
	}
	return m1
}

// 回文判断
// 字符串从左往右读和从右往左读都是一样的

func IsHuiWen(str string) (ret bool) {
	// 1、 len 不是字符的长度,所以应该用切片
	r := make([]rune, 0, len(str))
	for _, c := range str {
		r = append(r, c)
	}
	// 只需要判断到一半就可以了
	for i := 0; i < len(r)/2; i++ {
		if r[i] != r[len(r)-1-i] {
			return false
		}
	}
	return true
}
