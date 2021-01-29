package main

import "fmt"

// 强类型指定, 提前申明了一个变量res
// 返回值可以命名也可以不命名
func sum(x int, y int) (res int) {
	res = x + y
	return
}

// 无返回值函数
func f1(x int, y int) {
	fmt.Println(x + y)
}

// 没有参数没有返回值
func f2() {
	fmt.Println("f2")
}

// 没有参数但有返回值的
func f3() int {
	ret := 3
	return ret
}

// 多个返回值
func f5() (int, string) {
	return 1, "沙河"
}

// 可变长参数,
func f7(x string, y ...int) {
	fmt.Println(x)
	fmt.Println(y) // y是一个int类型的切片
}

// 参数类型简写
func f8(x, y, z int, k string) (int, string) {
	return x + y + z, k
}

// defer, 延迟执行, 多用于资源释放
func deferDemo() {
	fmt.Println("start")
	defer fmt.Println("defer-1")
	defer fmt.Println("defer-2")
	defer fmt.Println("defer-3")
	fmt.Println("end")
}

func main() {
	//fmt.Println(sum(2,3))
	//f7("2312",3,3,3,3,3)
	//fmt.Println(f8(1,2,3, "S"))
	deferDemo()
}
