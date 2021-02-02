package main

import "fmt"

type animal struct {
	name string
}

type dog struct {
	feet   uint8
	animal // 使用匿名嵌套
}

func (d dog) wang() {
	fmt.Println("狗在叫")
}

func (a animal) move() {
	fmt.Println("动物在动")
}

func main() {
	d1 := dog{
		feet: 4,
		animal: animal{
			name: "dog1",
		},
	}
	d1.wang()
	d1.move() // 使用匿名嵌套可以是实现嵌套访问
}
