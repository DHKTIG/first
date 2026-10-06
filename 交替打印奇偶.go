package main

import (
	"fmt"
	"sync"
)

func main() {
	// 定义两个无缓冲通道，用于互相通知
	chOdd := make(chan struct{})  // 控制奇数打印通道
	chEven := make(chan struct{}) // 控制偶数打印通道

	var wg sync.WaitGroup
	wg.Add(2) //定义一共有两个协程需要执行

	// 1. 打印奇数的协程
	go func() {
		defer wg.Done()
		for i := 1; i <= 10; i += 2 {
			<-chEven //将偶数协程中的空结构体取出来
			fmt.Println("奇数:", i)
			chOdd <- struct{}{} // 打印完后，塞一个空结构体到奇数协程的通道中，通知偶数协程可以打印了
		}
	}()

	// 2. 打印偶数的协程
	go func() {
		defer wg.Done()
		for i := 2; i <= 10; i += 2 {
			<-chOdd // 将奇数协程中的空结构体取出来
			fmt.Println("偶数:", i)
			chEven <- struct{}{} // 打印完后，塞一个空结构体到偶数协程的通道中，通知奇数协程可以打印了
		}
	}()

	// 3. 主协程发令：让奇数协程先起跑
	chEven <- struct{}{}

	wg.Wait() // 等待两个协程全部打印完毕
}
