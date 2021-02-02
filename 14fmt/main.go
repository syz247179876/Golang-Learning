package main

import "fmt"

func main() {
	fmt.Print("syz")

	fmt.Printf("%d%%\n\n", 12)
	fmt.Printf("%q\n", 65) // 转为ASCII
}
