package main

import "fmt"

func funcA() {
	// 立即释放资源
	defer func() {
		err := recover() // 尝试修复， recover必须搭配defer使用
		fmt.Println(err)
		fmt.Println("释放数据库连接！")
	}()
	panic("出错了") // 程序崩溃出错
	fmt.Println("a")
}

func funcB() {
	fmt.Println("b")
}

func main() {
	funcA()
	funcB()
}
