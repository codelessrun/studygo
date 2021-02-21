package main

import (
	"fmt"
	"time"
)

//时间
func f1() {
	now := time.Now()
	fmt.Println(now)
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())
	fmt.Println(now.Date())
	// 时间戳
	fmt.Println(now.Unix())     // 秒
	fmt.Println(now.UnixNano()) // 纳秒

	// 反向 由unix => time
	ret := time.Unix(1613709748, 0)
	fmt.Println(ret)

	// 时间间隔
	fmt.Printf("%d\n", time.Second)
	// now + 24小时
	fmt.Println(now.Add(24 * time.Hour))

	// sub 两个时间相减
	nextYear, err := time.Parse("2006-01-02", "2022-01-01")
	if err != nil {
		fmt.Printf("parse time failed, err: %v\n", err)
	}
	d := now.Sub(nextYear) // 求出时间间隔
	fmt.Println(d)
	fmt.Println("---------------")

	// 定时器
	//timer := time.Tick(time.Second)
	//for t := range timer{
	//	fmt.Println(t)  // 一秒钟执行一次
	//}

	//格式化时间
	// 2019-08-04
	fmt.Println(now.Format("2006-01-02"))
	// 2019/02/03 11:55:02
	fmt.Println(now.Format("2006/01/02 15:04:05"))
	// 2019/02/03 19:55:02 AM
	fmt.Println(now.Format("2006/01/02 03:04:05 PM"))
	// 2019/02/03 11:55:02.342
	fmt.Println(now.Format("2006/01/02 15:04:05.000"))
}

// 时区
func f2() {
	now := time.Now() // 本地的时间
	fmt.Println(now)
	// 明天的这个时间
	_, _ = time.Parse("2006-01-02 15:04:05", "2021-02-20 13:40:01")
	// 按照东八区的时区和格式解析一个字符串格式
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err!=nil {
		fmt.Printf("load loc failed, err: %v\n",err )
	}
	// 按照指定时区解析时间
	timeObj, err := time.ParseInLocation("2006-01-02 15:04:05", "2021-02-20 13:40:01", loc)
	if err!=nil {
		fmt.Printf("parse time failed, err: %v\n",err )
	}
	fmt.Println(timeObj)

	// 时间对象相减
	td := timeObj.Sub(now)
	fmt.Println(td)
}
func main() {
	//f1()
	f2()
}
