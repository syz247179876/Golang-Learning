package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name string `json:"name"` // 根据解析方式来重命名
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
}
