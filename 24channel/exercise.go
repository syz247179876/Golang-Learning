package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

var wg sync.WaitGroup

func getRand(jobs chan<- int64) {
	// 向chan中写入数据
	defer wg.Done() // 延迟调用
	for i := 0; i < 24; i++ {
		random := rand.Int63()
		fmt.Println("随机数字:", random)
		jobs <- random
	}
	close(jobs)
}

func compute(jobs <-chan int64, result chan<- int64) {
	// 从jobs中读出，写入到result
	defer wg.Done()
	for {
		i, ok := <-jobs
		if !ok {
			break
		}
		int64Num := i
		lenNum := len(strconv.FormatInt(int64Num, 10))
		var sum int64 = 0
		for i := 0; i < lenNum; i++ {
			last := int64Num % 10
			int64Num = int64Num / 10
			sum += last
		}
		result <- sum
	}
}

func main() {
	start := time.Now()
	jobsChan := make(chan int64, 100) // 使用有缓冲通道
	result := make(chan int64, 100)

	wg.Add(1)
	go getRand(jobsChan)

	for i := 0; i < 24; i++ {
		wg.Add(1)
		go compute(jobsChan, result)
	}
	wg.Wait()     // 阻塞直到24个goroutine都变成0
	close(result) // 要等到所有goroutine将结果计算完毕传给result通道后关闭通道
	for i := range result {
		fmt.Println(i)
	}
	end := time.Now()
	fmt.Println(end.Sub(start))

}
