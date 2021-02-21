package main

import (
	"fmt"
	"strings"
)

func main() {
	// 转义
	path := "\"D:\\gopath\\src\\code.oldboyedu.com\""
	fmt.Println(path)

	s := "I'm Ok"
	fmt.Println(s)

	// 反引号里的东西原样输出
	s2 := `
   春风又绿江两岸
   两岸怨声啼不知`
	fmt.Println(s2)
	s3 := `理想`
	fmt.Println(s3)
	fmt.Println(len(s3))

	// 分割
	ret := strings.Split(s3, "\\")
	fmt.Println(ret)
	// 包含
	fmt.Println(strings.Contains(s3, "理"))
	// 前缀
	fmt.Println(strings.HasPrefix(s3, "理"))
	// 后缀
	fmt.Println(strings.HasSuffix(s3, "想"))
	
	// indexOf 判断字串首次出现位置 ，lastIndexOf 判断最后一次出现的位置
	// join

}
