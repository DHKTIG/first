package main

import "fmt"

// 1. 生产者：源源不断产生数字
func generate() <-chan int {
	out := make(chan int)
	go func() {
		for i := 2; ; i++ { // 从2开始，无限循环
			out <- i
		}
	}()
	return out
}

// 检测是否为素数
func filter(in <-chan int, prime int) <-chan int {
	out := make(chan int)
	go func() {
		for num := range in {
			if num%prime != 0 { // 不能被当前素数整除，才放行
				out <- num
			}
		}
	}()
	return out
}

func main() {
	ch := generate() // 启动原料工厂

	fmt.Println("前10个素数：")
	for i := 0; i < 10; i++ {
		prime := <-ch // 从传送带上拿第一个数字，它必然是素数
		fmt.Println(prime)

		// 关键：基于这个素数，新建一个质检员，接替后面的流水线
		ch = filter(ch, prime)
	}
}
