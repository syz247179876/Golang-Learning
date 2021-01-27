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

	// 字符串修改
	s2 := "白萝卜"
	s3 := []rune(s2) // 把字符串转换成一个rune切片
	s3[0] = '红'
	fmt.Println(string(s3))

	// "" 和 ''类型比较
	c1 := "红"
	c2 := '红'
	fmt.Printf("c1:%T c2:%T\n", c1, c2) // c1:string  cs:int32（存储24bit的'红')

	// 类型转化
	n1 := 10
	var n2 float64
	n2 = float64(n1)
	fmt.Println(n2)

}
