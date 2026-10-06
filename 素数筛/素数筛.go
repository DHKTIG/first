package main

import "fmt"

// 生产者：源源不断产生数字
func generate() <-chan int {
	out := make(chan int)
	go func() {
		for i := 2; ; i++ { // 从2开始按顺序一直产生整数
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
			if num%prime != 0 { // 不能被当前素数整除放到传送带上
				out <- num
			}
		}
	}()
	return out
}

// 消费者：从传送带上拿数字，检测是否为素数
func main() {
	ch := generate() // 生产者

	fmt.Println("前10个素数：")
	for i := 0; i < 10; i++ {
		prime := <-ch // 从传送带上拿第一个数字，它必然是素数
		fmt.Println(prime)
		ch = filter(ch, prime) // 过滤掉能被当前素数整除的数字，剩下的数字继续传送
	}
}
