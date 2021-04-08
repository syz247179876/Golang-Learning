package main

import "fmt"

func main(){
	for i:=0; i< 10; i++{
		fmt.Println(i)
	}

	var i =5
	for ; i <10;i++{
		fmt.Println(i)
	}

	var j = 5
	for j< 10{
		fmt.Println(j)
		j++
	}
	
	s:= "你好云中"
	for i, v := range s{
		fmt.Printf("%d %c\n", i,v)
