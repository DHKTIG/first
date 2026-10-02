package main

import (
	"fmt"
	"sync"
)

func main() {
	// 定义两个无缓冲通道，用于互相通知
	chOdd := make(chan struct{})  // 奇数协程的接力棒
	chEven := make(chan struct{}) // 偶数协程的接力棒

	var wg sync.WaitGroup
	wg.Add(2) // 等待两个协程结束

	// 1. 打印奇数的协程
	go func() {
		defer wg.Done()
		for i := 1; i <= 10; i += 2 {
			<-chEven // 阻塞等待偶数协程传递过来的接力棒（一开始等主协程发令）
			fmt.Println("奇数:", i)
			chOdd <- struct{}{} // 打印完后，把接力棒传给偶数协程
		}
	}()

	// 2. 打印偶数的协程
	go func() {
		defer wg.Done()
		for i := 2; i <= 10; i += 2 {
			<-chOdd // 阻塞等待奇数协程传递过来的接力棒
			fmt.Println("偶数:", i)
			chEven <- struct{}{} // 打印完后，把接力棒传给奇数协程
		}
	}()

	// 3. 主协程发令：让奇数协程先起跑
	chEven <- struct{}{}

	wg.Wait() // 等待两个协程全部打印完毕
	fmt.Println("打印结束")
}
