package main

import (
	"fmt"
	"sync"
)

func main() {
	// 创建两个无缓冲通道用于控制打印顺序
	oddTurn := make(chan struct{})
	evenTurn := make(chan struct{})

	var wg sync.WaitGroup
	wg.Add(2) //两个 goroutine

	// 打印奇数
	go func() {
		defer wg.Done()
		for i := 1; i <= 10; i += 2 {
			<-evenTurn
			fmt.Println("奇数:", i)
			oddTurn <- struct{}{}
		}
	}()

	// 打印偶数
	go func() {
		defer wg.Done()
		for i := 2; i <= 10; i += 2 {
			<-oddTurn
			fmt.Println("偶数:", i)
			if i < 10 { // 避免在最后一次打印后再发送信号，导致死锁
				evenTurn <- struct{}{}
			}
		}
	}()

	// 让奇数先打印
	evenTurn <- struct{}{}

	wg.Wait()
}
