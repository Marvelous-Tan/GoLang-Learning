package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("=== Channel 修复版本 ===")
	
	// 使用WaitGroup来等待goroutine完成
	var wg sync.WaitGroup
	
	// 生产者消费者模式
	fmt.Println("生产者消费者模式:")
	ch := make(chan int, 5)
	
	// 生产者
	wg.Add(1)
	go func() {
		defer wg.Done()
		producer(ch)
	}()
	
	// 消费者
	wg.Add(1)
	go func() {
		defer wg.Done()
		consumer(ch)
	}()
	
	// 等待所有goroutine完成
	wg.Wait()
	fmt.Println("所有goroutine完成")
}

// 生产者函数
func producer(ch chan int) {
	for i := 1; i <= 5; i++ {
		ch <- i
		fmt.Printf("生产: %d\n", i)
		time.Sleep(100 * time.Millisecond)
	}
	close(ch) // 关闭channel
	fmt.Println("生产者完成")
}

// 消费者函数
func consumer(ch chan int) {
	for value := range ch { // 使用range接收直到channel关闭
		fmt.Printf("消费: %d\n", value)
		time.Sleep(200 * time.Millisecond)
	}
	fmt.Println("消费者完成")
} 