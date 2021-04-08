package main

import (
	"fmt"
	"sync"
)

// 并发安全的Map
var m = sync.Map{}

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(n int) {
			m.Store(n, n)
			value, _ := m.Load(n)
			fmt.Printf("k=:%v, v:=%v\n", n, value)
			wg.Done()
		}(i)

	}
	wg.Wait()
}
