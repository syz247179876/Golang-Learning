package main

import "fmt"

// GO语言中return不是原子操作
// 第一步:返回值赋值
// defer
// 第二步：真正的RET返回
// 函数中如果存在defer，则defer执行在第一步和第二步之间
func f1() int { // 没有指明返回值名字
	x := 5
	defer func() {
		x++
	}()
	return x // 为返回值5开辟一个空间
}

func f2() (x int) { // 为返回值x开辟一个空间
	defer func() {
		x++
	}()
	return 5 // 将x = 5, defer修改返回值x
}

func f3() (y int) { // 为返回值y开辟一个空间
	x := 5
	defer func() {
		x++
	}()
	return x // y = x = 5  , defer修改x
}
func f4() (x int) { // 为x开辟一个空间
	defer func(x int) {
		x++ // 修改x的副本
	}(x)
	return 5 // x = 5
}

// 传递x的指针
func f5() (x int) { // 为x开辟一个空间
	defer func(x *int) {
		*x++ // 修改地址指向的内存数据
	}(&x) // 传递地址
	return 5 // x = 5
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

// defer嵌套
func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func main() {
	//fmt.Println(f1()) // 5
	//fmt.Println(f2()) // 6
	//fmt.Println(f3()) // 5
	//fmt.Println(f4()) // 5
	//fmt.Println(f5()) // 6

	//fmt.Printf("%T\n", f1) // func() int
	//f6(f2)                 // 6
	//fk := f7(f2)
	//fk(3, 3) // ff

	x := 1
	y := 2
	defer calc("AA", x, calc("A", x, y)) // 将x:=1,y:=2压入栈
	x = 10
	defer calc("BB", x, calc("B", x, y))
	y = 20
}
