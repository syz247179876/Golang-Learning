package main

import (
	"fmt"
	"sync"
)

type singleton struct {
}

var instance *singleton
var once sync.Once

// 并发安全的单例模式
func getInstance() *singleton {
	// sync.Once内部包含了一个互斥锁和一个布尔值
	// 互斥锁用来保证布尔值和数据的安全
	// 布尔值用来记录初始化是否完成
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}

// 非单例模式+非并发安全
func createInstance() *singleton {
	return &singleton{}
}

func main() {
	var instance1 = getInstance()
	var instance2 = getInstance()
	var instance3 = createInstance()
	var instance4 = createInstance()
	fmt.Println(instance3 == instance4) // false
	fmt.Println(instance1 == instance2) // true
}
