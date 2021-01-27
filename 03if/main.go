package main

import "fmt"

func main() {
	age := 19
	if age > 18 {
		fmt.Println("哦里干")
	}
	// age作为局部变量，只在循环体内执行
	if age := 19; age > 18 {
		fmt.Println("做羞羞的事情")
	}
}
