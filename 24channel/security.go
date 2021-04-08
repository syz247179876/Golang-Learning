package main

import (
	"fmt"
	"sync"
	"time"
)

/* golang中并发安全和锁机制 */

var wgs sync.WaitGroup
var lock sync.Mutex
var rwlock sync.RWMutex
var x int64

func addWithNoLock() {
	for i := 0; i < 100; i++ {
		x += 1
	}
	wgs.Done() // goroutine执行完毕-1
}

func addWithLock() {
	for i := 0; i < 100; i++ {
		lock.Lock() // 加锁
		x += 1
		fmt.Println(x)
		lock.Unlock() // 解锁
	}
	wgs.Done()
}

// 写操作
func write() {
	rwlock.Lock() // 加上写锁
	x = x + 1
	time.Sleep(10 * time.Millisecond) // 模拟写耗时10ms
	rwlock.Unlock()                   // 解写锁
	wgs.Done()
}

// 读操作
func reader() {
	rwlock.RLock() // 加读锁
	//time.Sleep(time.Millisecond) // 模拟读耗时1ms
	rwlock.RUnlock() // 解读锁
	wgs.Done()
}

func main() {
	//wgs.Add(2)  // 启动2个goroutine+2
	//go addWithNoLock()
	//go addWithNoLock()

	//go addWithLock()
	//go addWithLock()
	start := time.Now()
	for i := 0; i < 10; i++ {
		wgs.Add(1)
		go write()
	}

	for i := 0; i < 1000; i++ {
		wgs.Add(1)
		go reader()
	}

	wgs.Wait() // 等待所有的goroutine执行完毕再往下执行
	end := time.Now()

	fmt.Println(end.Sub(start))
}
