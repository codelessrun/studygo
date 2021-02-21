package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

// 文件写

func writeDemo1() {
	fileObj, err := os.OpenFile("./xx.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Printf("open file failed , err :%v", err)
		return
	}
	defer fileObj.Close()
	// write
	fileObj.Write([]byte("zhoulin mengbi le\n"))
	// write
	fileObj.WriteString("周林解释不了了")
}

func writeDemo2() {
	fileObj, err := os.OpenFile("./xx.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Printf("open file failed , err :%v", err)
		return
	}
	defer fileObj.Close()
	//	创建一个写的对象
	wr := bufio.NewWriter(fileObj)
	_, _ = wr.WriteString("hello沙河\n") // 写到缓存中
	_ = wr.Flush()                     // 将缓存中的文件写入实际文件中
}

func writeDemo3()  {
	str := "hello 沙河你好啊"
	err := ioutil.WriteFile("./xx.txt",[]byte(str),0666)
	if err!=nil {
		fmt.Printf("write file failed,err %v",err)
		return
	}
}

func main() {
	//writeDemo1()
	//writeDemo2()
	writeDemo3()
}
