package main

import "fmt"

func main() {

	// 函数内部生命不带名字的函数
	f1 := func(x, y int) {
		fmt.Println(x + y)
	}
	f1(10, 20)

	// 只执行一次，立即执行函数
	func(x, y int) {
		fmt.Println(x + y)
	}(10, 20)

}
