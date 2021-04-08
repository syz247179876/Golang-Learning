package main

import "fmt"

// 常量
const pi = 3.1415

// 批量申明常量
const (
	statusOK = 200
	notFOUND = 404
)
const (
	name = "syz"
	age = 19
)

// 如果某一行申明后没有赋值，默认和上一行一致
const (
	n1 = 100
	n2
	n3
)

// iota
const (
	a1 = iota // 0
	a2        // 1
	a3        // 2
	a4        // 3
)

// 匿名
const (
	b1 = iota // 0
	b2        // 1
	_         // 2, 丢弃
	b3        // 3
)

// 插队
const (
	c1 = iota // 0
	c2 = 100  // 100
	c3 = iota // 2
	c4
)

// 多个常量申明在一行
const (
	d1, d2 = iota + 1, iota + 2 // d1:1 d2:2
	d3, d4 = iota + 1, iota + 3 // d3:2 d4:4
)

// 定义数量级
const (
	_  = iota // 丢掉
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
)

func define_fun(){
	n := 23
	m := 200
	fmt.Println(n, m)
}


func main() {
	define_fun()
	fmt.Println("a1", a1)
	fmt.Println("a2", a2)
	fmt.Println("a3", a3)
	fmt.Println("b1", b1)
	fmt.Println("b2", b2)
	fmt.Println("b3", b3)
	fmt.Println("c1", c1)
	fmt.Println("c2", c2)
	fmt.Println("c3", c3)
	fmt.Println("c4", c4)
	fmt.Println("d1", d1)
	fmt.Println("d2", d2)
	fmt.Println("d3", d3)
	fmt.Println("d4", d4)

}
