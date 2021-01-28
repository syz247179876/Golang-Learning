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

	w1 := make([]int, 5, 10)
	fmt.Printf("w1=%v len(w1)=%d cap(s1)=%d\n", w1, len(w1), cap(w1))

	// 切片遍历
	for i := 0; i < len(s2); i++ {
		fmt.Println(s2[i])
	}
	for i, v := range s2 {
		fmt.Println(i, v)
	}

	// 追加元素,进行扩容
	z1 := []string{"江苏", "上海", "深圳"}
	fmt.Printf("z1=%v len(z1)=%d cap(z1)=%d\n", z1, len(z1), len(z1))
	z1 = append(z1, "广州") // 要用原来的切片
	fmt.Printf("z1=%v len(z1)=%d cap(z1)=%d\n", z1, len(z1), len(z1))

	z1 = append(z1, "杭州", "南京")

	ss := []string{"武汉", "成都"}
	z1 = append(z1, ss...) // ...表示拆开
	fmt.Printf("z1=%v len(z1)=%d cap(z1)=%d\n", z1, len(z1), len(z1))

	// copy

	k1 := []int{1, 3, 5}
	k2 := k1 // 赋值， 相当于浅拷贝
	k3 := make([]int, 3, 3)
	copy(k3, k1) // 相当于深拷贝
	fmt.Println(k1, k2, k3)
	k1[0] = 2
	fmt.Println(k1, k2, k3)

	// 删除切片元素, 指针指向
	x1 := [...]int{1, 3, 5, 7, 9, 11}
	x2 := x1[:] // 构造一个切片
	x2 = append(x2[:1], x2[4:]...)
	fmt.Println(x1)

	var e = make([]int, 5, 10)
	for i := 0; i < 10; i++ {
		e = append(e, i)
	}
	fmt.Println(e)

}
