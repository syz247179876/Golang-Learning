package main

import (
	"fmt"
	"os"
)

func showMenu() {
	fmt.Println("go版本学生管理系统")
	fmt.Println(`
     1. 查看所有学生
     2. 添加学生
     3. 修改学生
     4. 删除学生
     5.退出
    `)
}

var m manager

func main() {
	m = manager{
		allStudent: make(map[int64]student, 100), // 开辟空间
	}
	for {
		showMenu()
		fmt.Print("请输入选项:")
		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			m.showStu()
		case 2:
			m.addStu()
		case 3:
			m.modifyStu()
		case 4:
			m.deleteStu()
		case 5:
			os.Exit(1)
		default:

		}
	}

}
