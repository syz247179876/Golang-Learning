package main

import "fmt"

// GO语言中return不是原子操作
// 第一步:返回值赋值
// defer
// 第二步：真正的RET返回
// 函数中如果存在defer，则defer执行在第一步和第二步之间
func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
}
func f4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5
}

// 函数作为参数传入
func f6(x func() int) {
	ret := x()
	fmt.Println(ret)
}

func ff(x, y int) int {
	fmt.Println("ff")
	return 2
}

// 函数作为返回值
func f7(x func() int) func(int, int) int {
	return ff
}

func main() {
	fmt.Println(f1()) // 5
	fmt.Println(f2()) // 6
	fmt.Println(f3()) // 5
	fmt.Println(f4()) // 5

	fmt.Printf("%T\n", f1) // func() int
	f6(f2)                 // 6
	fk := f7(f2)
	fk(3, 3) // ff
}
