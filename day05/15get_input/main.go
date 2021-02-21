package main

import (
	"bufio"
	"fmt"
	"os"
)

// 获取用户输入时，如果遇到空格

func userScan()  {
	var s string
	fmt.Print("请输入内容:")
	_, _ = fmt.Scanln(&s) // 这个方法是读到空白符停止。空白符包括空格换行
	fmt.Printf("你输入的内容是:%s\n",s)
}

func userBufio()  {
	var s string
	fmt.Print("请输入内容:")
	reader :=bufio.NewReader(os.Stdin) //从标准输入去读
	s, _ = reader.ReadString('\n') // 读到换行符就不读了
	fmt.Printf("你输入的内容是:%s\n",s)
}

func main() {
	//userScan()
	userBufio()
}
