package main

import "fmt"

func main() {
	// 基于数组定义切片，没有长度限制
	var s1 []int
	var s2 []string
	fmt.Println(s1 == nil) // 没有开辟内存空间， true

	s1 = []int{1, 2, 3}
	s2 = []string{"syz", "zjw"}

	fmt.Println(s1, s2)

	// 由数组得到切片
	a1 := [...]int{1, 3, 5, 7, 9, 11}
	s3 := a1[0:4]
	s4 := a1[2:]
	fmt.Println(s3)

	fmt.Println(a1)
	// 切片是引用类型
	a1[2] = 300

	// cap表示底层数组从切片的第一个元素到最后的元素的数量
	fmt.Printf("len(s3):%d cap(s3):%d\n", len(s3), cap(s3))
	fmt.Printf("len(s4):%d cap(s4):%d\n", len(s4), cap(s4))

	// 切片在切片
	s5 := s4[2:]
	fmt.Printf("len(s5):%d cap(s5):%d\n", len(s5), cap(s5))
	fmt.Println(a1)

	// make函数创建切片

}
