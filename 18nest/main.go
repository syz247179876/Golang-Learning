package main

import "fmt"

type address struct {
	province string
	city     string
}

type address2 struct {
	city string
}

type person struct {
	name string
	address
	address2
}

func main() {
	p1 := person{
		name: "syz",
		address: address{
			province: "江苏",
			city:     "扬州",
		},
		address2: address2{
			city: "杭州",
		},
	}

	//fmt.Println(p1.city)  // 先在person结构体中查找，找不到再去匿名嵌套结构体中查找
	// 如果匿名嵌套结构体冲突，则需显示指定哪个结构体的字段
	fmt.Println(p1.address2.city)
}
