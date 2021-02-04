package main

import "fmt"

type student struct {
	id   int64
	name string
}

type manager struct {
	allStudent map[int64]student
}

// 查看学生信息
func (s manager) showStu() {
	// 遍历map
	for _, v := range s.allStudent {
		fmt.Printf("学号:%d 姓名:%s\n", v.id, v.name)
	}
}

// 增加学生
func (s manager) addStu() {
	var ( // 定义变量
		stuId   int64
		stuName string
	)
	fmt.Print("请输入学号:")
	fmt.Scanln(&stuId)
	fmt.Print("请输入姓名:")
	fmt.Scanln(&stuName)
	newStudent := student{
		id:   stuId,
		name: stuName,
	}
	s.allStudent[newStudent.id] = newStudent
	fmt.Println("添加学生")
}

// 删除学生
func (s manager) deleteStu() {

	var stuID int64
	fmt.Print("请输入学号:")
	fmt.Scanln(&stuID)
	_, ok := s.allStudent[stuID]
	if !ok {
		fmt.Println("不存在该学号")
	} else {
		delete(s.allStudent, stuID) // 从map中删除键值对
		fmt.Println("删除成功")
	}

}

// 修改学生
func (s manager) modifyStu() {
	var stuID int64
	fmt.Print("请输入学号:")
	fmt.Scanln(&stuID)

	stuObj, ok := s.allStudent[stuID]
	if !ok {
		fmt.Println("不存在该学号")
	} else {
		fmt.Print("请输入学生新名字:")
		var newName string
		fmt.Scanln(&newName)
		stuObj.name = newName
		s.allStudent[stuID] = stuObj
		fmt.Println("修改完成")
	}

}
