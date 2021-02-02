package main

import "fmt"

type dog struct {
	name string
	age  int
}

// 方法是作用于特定类型的函数，指明谁来接受
// 接受者表示调用该方法的具体类型变量，用类型名小写字母表示
// 传的是拷贝
func (d dog) wang() {
	fmt.Printf("%s:旺旺", d.name)
}

// 传的是指针
func (d *dog) sleep() {
	d.age++
}

func newDog(name string, age int) dog {
	return dog{
		name: name,
		age:  age,
	}
}

func main() {
	dog1 := newDog("dog1", 19)
	dog1.wang()
	dog1.sleep()
	fmt.Println(dog1.age)
}
