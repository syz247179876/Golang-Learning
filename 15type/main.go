package main

import "fmt"

// 自定义类型

// type后面跟的类型

type myInt int     // 自定义类型
type yourInt = int // 类型别名

func main() {
	var n myInt
	var m yourInt
	n = 100
	m = 50
	fmt.Println(n)
	fmt.Printf("%T\n", n)
	fmt.Println(m)
	fmt.Printf("%T\n", m)

	var c rune // 存放字符
	c = '四'
	fmt.Printf("%v", c)
}
