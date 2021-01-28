package main

import "fmt"

func main() {
	var m1 map[string]int
	m1 = make(map[string]int, 10)
	m1["syz"] = 20
	m1["zjw"] = 21

	fmt.Println(m1)
	fmt.Println(m1["syz"])
	// 用ok来接受bool值，如果接受不存在的值就会返回零值
	value, ok := m1["lsp"]
	if !ok {
		fmt.Println("无key")
	} else {
		fmt.Println(value)
	}

	for k, v := range m1 {
		fmt.Println(k, v)
	}

	// 遍历key
	for k := range m1 {
		fmt.Println(k)
	}
	// 只遍历value
	for _, value := range m1 {
		fmt.Println(value)

	}

	delete(m1, "syz")
	fmt.Println(m1)

	// 元素类型为map的切片，map和map组合
	s1 := make([]map[int]string, 10, 10)
	s1[0] = make(map[int]string, 1) // 对切片进行初始化
	s1[0][10] = "syz"
	fmt.Println(s1)

	// 值为切片类型的map

	m2 := make(map[string][]int, 10)
	m2["name"] = []int{10, 20, 30} // 对切片进行初始化
	fmt.Println(m2)
}
