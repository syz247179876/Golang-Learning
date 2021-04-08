package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name string `json:"name"` // 根据解析方式来重命名 // 大些变量能够被外部引用
	Age  int    `json:"age"`
}

func main() {
	p1 := person{
		Name: "syz",
		Age:  20,
	}
	fmt.Println(p1)

	b, err := json.Marshal(p1) // 序列化
	if err != nil {            // 如果序列化成功
		fmt.Println(err)
	}
	fmt.Println(string(b)) // 转为字符串
	var p2 person
	str := `{"name":"理想","age":18}`
	json.Unmarshal([]byte(str), &p2)  // 传指针修改p2的值，而不是拷贝
	fmt.Println(p2)
}
