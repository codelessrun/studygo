package main

import "fmt"

// 学生管理系统

type student struct {
	id   int64
	name string
}

type studentMgr struct {
	allStudent map[int64]*student
}

// 查看学生
func (s studentMgr) showStudent() {
	for _, stu := range s.allStudent {
		fmt.Printf("学号: %d 姓名: %s\n", stu.id, stu.name)
	}
}

// 增加学生
func (s studentMgr) addStudent() {
	var (
		stuId int64
		name  string
	)

	// 获取用户输入
	fmt.Print("请输入学号:> ")
	_, _ = fmt.Scanln(&stuId)
	if stuId==0 {
		// 因为学号应该是数字，但是如果用户输入的是字符，则学号会被当作0,所以这里限制不能为0
		fmt.Println("学号不能是0或者字符")
		return
	}
	fmt.Print("请输入姓名:> ")
	_, _ = fmt.Scanln(&name)



	// 根据用户输入创建对象
	newStu := student{
		id:   stuId,
		name: name,
	}

	//	2、把新的学生放到 s.allStudent中
	s.allStudent[newStu.id] = &newStu
	fmt.Println("添加成功!")
}

// 编辑学生
func (s studentMgr) editStudent() {
	// 1、获取用户输入的学号
	var stuId int64
	fmt.Print("请输入要修改的学号:> ")
	_, _ = fmt.Scanln(&stuId)
	if stuId==0 {
		// 因为学号应该是数字，但是如果用户输入的是字符，则学号会被当作0,所以这里限制不能为0
		fmt.Println("学号不能是0或者字符")
		return
	}

	// 2、展示对应的学生信息，如果没有显示查无此人
	//var student *student
	//for _, stu := range s.allStudent {
	//	if stu.id == stuId {
	//		student = stu
	//	}
	//}
	//if student != nil {
	//	// 3、请输入修改后的姓名
	//	fmt.Printf("要修改的学生信息: 学号:%d 姓名: %s\n", student.id, student.name)
	//	fmt.Print("请输入修改后的姓名:> ")
	//	_, _ = fmt.Scanln(&student.name)
	//
	//	// 4、更新学生姓名
	//	s.allStudent[student.id] = student
	//	fmt.Println("修改成功！")
	//} else {
	//	fmt.Println("查无此人!")
	//}

	//2、展示对应的学生信息，如果没有显示查无此人
	student, ok := s.allStudent[stuId]
	if !ok {
		fmt.Println("查无此人！")
		return
	}
	fmt.Printf("你要修改的学生信息如下: 学号: %d 姓名: %s\n", student.id, student.name)

	// 3、请输入修改后的姓名
	fmt.Print("请输入修改后的姓名:> ")
	_, _ = fmt.Scanln(&student.name)

	//4、更新学生姓名
	s.allStudent[student.id] = student
	fmt.Println("修改成功！")
}

// 删除学生
func (s studentMgr) deleteStudent() {
	//	1、获取用户输入学号
	var stuId int64
	fmt.Print("请输入要删除的学号:> ")
	_, _ = fmt.Scanln(&stuId)
	if stuId==0 {
		// 因为学号应该是数字，但是如果用户输入的是字符，则学号会被当作0,所以这里限制不能为0
		fmt.Println("学号不能是0或者字符")
		return
	}

	//	2、去找这个学生，如果没有则显示查无此人
	_, ok := s.allStudent[stuId]
	if !ok {
		fmt.Println("查无此人！")
		return
	}

	//  3、有的话就删除
	delete(s.allStudent, stuId)
	fmt.Println("删除成功！")
}
