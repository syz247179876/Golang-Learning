package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

// 创建channel
func createChan() {

	var channel1 chan int // 声明一个传递整型的通道
	//var channel2 chan []int // 声明一个int切片的通道
	//var channel3 chan int

	channel1 = make(chan int) // channel是引用类型，必须初始化才能使用
	//channel2 = make(chan []int)
	//channel3 = make(chan int)

	channel1 <- 10  // 将10发送到channel1中
	x := <-channel1 // 从channel1中接受值并赋值给变量x

	fmt.Println(x)
	fmt.Println("你好")
	close(channel1)

	//fmt.Println(channel2, channel1, channel3)
}

func rangeFor() {
	// 开启一个goroutine定义一个匿名函数，发送值到ch1
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			ch1 <- i // 向ch3中添加100个数
		}
		close(ch1)
	}()

	// 开启goroutine从ch1中接受值， 并将该值的平方发送到ch2中
	go func() {
		for {
			i, ok := <-ch1 // 通道关闭后，ok为false
			if !ok {
				break
			}
			ch2 <- i * i
		}
		close(ch2)
	}()

	// 开启goroutine从ch2中接受值打印
	for i := range ch2 { // 如果通道关闭，则会退出for range操作
		fmt.Println(i)
	}
}

func recv(c chan int) { // c为双向通道
	ret := <-c // 如果没有发送值过来则会阻塞
	fmt.Println("接受成功", ret)
}

// 写通道
func writer(in chan<- int) {
	for i := 1; i < 100; i++ {
		in <- i
	}
	close(in)
}

// 读写通道
func wr(in chan<- int, out <-chan int) {
	for i := range out {
		in <- i * i
	}
	close(in)
}

// 读通道
func read(out <-chan int) {
	for i := range out {
		fmt.Println(i)
	}
}

func main() {
	ch := make(chan int)     // 定义一个无缓冲通道,即同步通道
	ch2 := make(chan int, 1) // 创建一个容量为1的有缓冲区通道
	go recv(ch)
	ch <- 10
	go recv(ch2)
	fmt.Println("发送成功")

	/* -----------------for range------------------- */

	// rangeFor()  // 使用for range 遍历channel

	/* ----------------单向通道---------------------- */
	inner := make(chan int)
	outer := make(chan int)
	go writer(inner)    // 向通道中写数据
	go wr(outer, inner) // 遍历inner，写入outer
	read(outer)         // 遍历outer
}
