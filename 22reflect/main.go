package main

import (
	"fmt"
	"reflect"
)

// 通过反射机制获取类型
func reflectType(x interface{}) {
	v := reflect.TypeOf(x)
	fmt.Printf("type:%v\n", v)
}
func getType() {
	var a float32 = 3.14
	reflectType(a)

	var b int64 = 100
	reflectType(b)
}

// 自定义类型
type myInt int64

// 获取接口的类型和种类
func reflectType1(x interface{}) {
	t := reflect.TypeOf(x)
	fmt.Printf("type:%v kind:%v\n", t.Name(), t.Kind())
}

// 比较type和kind
func getKind() {
	var a *float32
	var b myInt
	var c rune // 类型别名
	reflectType1(a)
	reflectType1(b)
	reflectType1(c)

	type Person struct {
		name string
		age  int
	}

	type book struct{ title string }

	var d = Person{
		name: "司云中",
		age:  22,
	}
	book1 := book{
		title: "悲伤逆流成河",
	}
	reflectType1(d)
	reflectType1(book1)

}

// 通过反射修改值，由于函数参数传递的是值拷贝
func reflectSetValue2(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Kind() == reflect.Int64 {
		v.SetInt(200) // 修改的是副本， reflect包含引发panic
	}
}

func reflectSetValue3(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Elem().Kind() == reflect.Int64 {
		v.Elem().SetInt(200) // 使用Elem来获取指针的值，然后进行修改
	}
}
func isNil() {
	// *int类型空指针
	var a *int
	fmt.Println("var a *int IsNil:", reflect.ValueOf(a).IsNil())
	// nil值
	fmt.Println("nil IsValid:", reflect.ValueOf(nil).IsValid())
	// 实例化一个匿名结构体
	b := struct{}{}
	// 尝试从结构体中查找"abc"字段
	fmt.Println("不存在的结构体成员:", reflect.ValueOf(b).FieldByName("abc").IsValid())
	// 尝试从结构体中查找"abc"方法
	fmt.Println("不存在的结构体方法:", reflect.ValueOf(b).MethodByName("abc").IsValid())
	// map
	c := map[string]int{}
	// 尝试从map中查找一个不存在的键
	fmt.Println("map中不存在的键：", reflect.ValueOf(c).MapIndex(reflect.ValueOf("娜扎")).IsValid())
}
func main() {
	getType()
	getKind()

	var x int64 = 777

	// reflectSetValue2(x)
	reflectSetValue3(&x)
	fmt.Println(x)

	isNil()

}
