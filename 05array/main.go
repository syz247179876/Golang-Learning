package main

import "fmt"

func main() {
	var a1 [3]bool
	var a2 [4]bool
	fmt.Println(a1, a2)
	a1 = [3]bool{false, false, false}
	// ... 表示自动推断数组的长度是多少
	a100 := [...]int{0, 12, 3, 3, 41, 2, 3, 1, 23, 123, 2, 31}
	fmt.Println(a100)
	// 根据索引进行初始化
	a4 := [5]int{0: 5, 2: 4, 4: 1}
	fmt.Println(a4)

	// 数组的遍历
	city := [...]string{"北京", "上海", "深圳"}
	// 常规遍历
	for i := 0; i < len(city); i++ {
		fmt.Println(city[i])
	}
	// for range

	for i, v := range city {
		fmt.Println(i, v)
	}

	// 多维数组
	var b [3][2]int
	b = [3][2]int{
		[2]int{1, 2},
		[2]int{3, 4},
		[2]int{5, 6},
	}
	fmt.Println(b)

	for _, v := range b {
		for j, v2 := range v {
			fmt.Println(j, v2)
		}
	}

	// 数组值类型,直接存储数据，定义后就为其分配内存
	b1 := [3]int{1, 2, 3}
	b2 := b1
	b2[0] = 100
	for _, v := range b1 {
		fmt.Println(v)
	}

	dd := [...]int{1, 3, 5, 7, 8}
	n := len(dd)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if dd[i]+dd[j] == 8 {
				fmt.Printf("(%d,%d)", i, j)
			}
		}
	}
}
