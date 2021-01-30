package main

import "fmt"

// 定义的类型名和字段名不能重复
type person struct {
	name   string
	age    int
	gender string
	hobby  []string
}

func main() {
	var syz person

	// 通过字段赋值
	syz.name = "syz"
	syz.age = 9000
	syz.hobby = []string{"睡觉", "吃饭"}
	syz.gender = "男"
	fmt.Println(syz)
	fmt.Println(syz.age)

	// 匿名结构体
	var s struct {
		x string
		y int
	}

	s.x = "嘿嘿嘿"
	s.y = 100
	fmt.Printf("type:%T value:%v", s, s)

}
