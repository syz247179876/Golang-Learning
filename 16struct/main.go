package main

import "fmt"

// 定义的类型名和字段名不能重复
type person struct {
	name   string
	age    int
	gender string
	hobby  []string
}

type person1 struct {
	name, gender string
}

// go语言中函数参数永远是拷贝
func f(x *person1) {
	(*x).gender = "女" // 根据内存地址修改
	x.gender = "男"    // 语法糖形式
}

// struct结构体中定义的数据内存地址是连续的
type test struct {
	a int8
	b int8
	c int8
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

	var zjw person1

	f(&zjw)
	fmt.Println(zjw)

	var p1 = new(person)
	fmt.Println(p1)
	p1.gender = "男"
	(*p1).gender = "女"
	fmt.Println(p1)
	fmt.Printf("%p\n", p1)  // p1保存的内存地址
	fmt.Printf("%p\n", &p1) //  p1的内存地址

	// 初始化并复制
	var p3 = person{
		name:   "222",
		gender: "男",
	}
	fmt.Printf("%v", p3)

	var p4 = test{
		a: 100,
		b: 100,
		c: 100,
	}
	fmt.Printf("%p\n", &p4.a)
	fmt.Printf("%p\n", &p4.b)
	fmt.Printf("%p\n", &p4.c)

}
