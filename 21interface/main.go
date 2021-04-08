package main

import "fmt"

// 只在乎行为，不在乎类型
type speaker interface {
	speak() // 只要实现了speak，就将其当作speaker类型
}

type cat struct {
}

type dog struct {
}

func (c cat) speak() {
	fmt.Println("猫叫")
}

func (d dog) speak() {
	fmt.Println("狗叫")
}

func da(x speaker) {
	x.speak()
}

type Personer interface {
	write()
}

type student struct{}
type teacher struct{}

// 指针接受着实现接口
func (stu *student) write() {
	fmt.Printf("指针接受者实现接口")
}

// func (stu student) write() {
// 	fmt.Printf("学生写作")
// }
func (tea teacher) write() {
	fmt.Printf("老师写作")
}

func calling() {
	var p Personer
	s := &student{} // 实例化一个student, 指针接受者实现
	t := teacher{}  // 实例化一个dog
	p = s
	p.write()
	p = t
	t.write()
}

type WashingMachine interface {
	wash()
	dry()
}

type dryer struct{}

func (d dryer) dry() {
	fmt.Printf("甩干")
}

type haier struct {
	dryer
}

func (h haier) wash() {
	fmt.Printf("海尔洗衣机")
}

func washs() {
	var w WashingMachine
	h := haier{}
	w = h
	w.wash()
	w.dry()
}

func nullinterface() {
	var x interface{} // 定义一个空接口
	s := "司云中"
	x = s
	fmt.Printf("type:%T, value:%v", x, x)
}

// 空姐口作为map值实现任意值的字典
func nullMap() {
	var studentInfo = make(map[string]interface{})
	studentInfo["name"] = "司云中"
	studentInfo["age"] = 21
	studentInfo["married"] = false
	fmt.Println(studentInfo)
}

// 判断接口中值的类型
func judgeType() {
	var x interface{}
	x = "司云中"
	v, ok := x.(string)
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("断言失败")
	}
}

func main() {
	var c1 cat
	var d1 dog
	da(c1)
	da(d1)
	calling()
	washs()
	// nullinterface()
	nullMap()
	judgeType()
}
