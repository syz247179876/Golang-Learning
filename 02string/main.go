package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "I'm ok"
	fmt.Println(s)

	// 字符串长度
	fmt.Println(len(s))

	// 字符串凭借
	name := "理想"
	world := "丰满"
	ss := name + world
	fmt.Println(ss)
	// 字符串拼接
	ss1 := fmt.Sprintf("%s%s", name, world)
	fmt.Println(ss1)
	// 字符串分割
	ret := strings.Split(s, "o")
	fmt.Println(ret)
	// 包含
	fmt.Println(strings.Contains(s, "ok"))

	// 前缀和后缀
	fmt.Println(strings.HasPrefix(ss, "理想"))
	fmt.Println(strings.HasSuffix(ss, "不丰满"))

	// 下标
	s4 := "syzzjw.cn"
	fmt.Println(strings.Index(s4, "j"))

	// 拼接
	fmt.Println(strings.Join(ret, "+"))
}
