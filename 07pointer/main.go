package main

import "fmt"

func main() {
	// 1.&
	// 2.*
	n := 18
	fmt.Println(&n)

	p := &n
	fmt.Println(*p)

	// 引用类型只在栈上为地址分配内存，并未在堆上为数据分配内存

	var a *int
	a = new(int) // 在堆上为a开辟内存
	*a = 100
	fmt.Println(*a)

}
