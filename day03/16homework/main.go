package main

import (
	"fmt"
	"strings"
)

/**
分金币
有一堆人
给这些人按名字中字母不同分不同金币
求剩下的金币
*/
var (
	coins        = 50
	users        = []string{"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth"}
	distribution = make(map[string]int, len(users))
)

func main() {
	//left := dispatchCoin()
	left := dispatchCoin2()
	fmt.Println("剩下的: ", left)
	fmt.Println(distribution)
}

// 我写的
func dispatchCoin() (left int) {
	sum := 0 // 总共分掉多少金币
	for _, name := range users {
		count := 0
		// 计算这个人应该分多少金币
		for _, r := range name {
			v := string(r)
			switch strings.ToUpper(v) {
			case "E":
				count++
			case "I":
				count += 2
			case "O":
				count += 3
			case "U":
				count += 4
			}
		}
		distribution[name] = count
		sum += count
	}
	left = coins - sum
	return
}

func dispatchCoin2() (left int) {
	// 1. 依次拿到每个人到名字
	for _, name := range users {
		// 2. 拿到一个人人名,根据分金币的规则取分金币
		for _, c := range name {
			switch c {
			case 'e', 'E':
				distribution[name]++
				coins--
			case 'i', 'I':
				distribution[name] += 2
				coins -= 2
			case 'o', 'O':
				distribution[name] += 3
				coins -= 3
			case 'u', 'U':
				distribution[name] += 4
				coins -= 4
			}
		}
	}
	left = coins
	return
}
