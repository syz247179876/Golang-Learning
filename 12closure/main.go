package main

import "fmt"

// 闭包

func f1(f func()) {
	f()
}

func f2(x, y int) {
	fmt.Println(x + y)
}

func f3(f func(int, int), m, n int) func() {
	temp := func() {
		fmt.Println("hello")
		f(m, n)
	}
	return temp
}

// 构造一个新的函数实现f1中调用f2
func f4(f func(int, int), m, n int) func() {
	temp := func() {
		f(m, n)
	}
	return temp
}

// 闭包引用
func calculate(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}
	sub := func(i int) int {
		base -= i
		return base
	}
	return add, sub
}

func main() {
	//f3(f2, 3,4)()
	f1(f4(f2, 4, 3)) //7

	f1, f2 := calculate(10)
	fmt.Println(f1(1), f2(3)) // 11, 8
	fmt.Println(f1(3), f2(4)) // 11, 7

}
