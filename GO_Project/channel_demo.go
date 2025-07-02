package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("=== Channel 演示 ===")
	
	// 1. 创建 channel
	ch := make(chan int) // 无缓冲channel
	ch2 := make(chan string, 3) // 有缓冲channel，容量为3
	
	// 2. 基本的发送和接收
	fmt.Println("1. 基本发送接收:")
	go func() {
		ch <- 42 // 发送数据到channel
		fmt.Println("数据已发送到channel")
	}()
	
	value := <-ch // 从channel接收数据
	fmt.Printf("从channel接收到: %d\n", value)
	
	// 3. 有缓冲channel
	fmt.Println("\n2. 有缓冲channel:")
	ch2 <- "消息1"
	ch2 <- "消息2"
	ch2 <- "消息3"
	
	fmt.Println(<-ch2) // 接收消息1
	fmt.Println(<-ch2) // 接收消息2
	fmt.Println(<-ch2) // 接收消息3
	
	// 4. 生产者消费者模式
	fmt.Println("\n3. 生产者消费者模式:")
	ch3 := make(chan int, 5)
	
	// 生产者
	go producer(ch3)
	
	// 消费者
	go consumer(ch3)
	
	// 等待一段时间
	time.Sleep(3 * time.Second)
	
	// 5. select语句
	fmt.Println("\n4. Select语句演示:")
	ch4 := make(chan string)
	ch5 := make(chan string)
	
	go func() {
		time.Sleep(1 * time.Second)
		ch4 <- "来自channel1"
	}()
	
	go func() {
		time.Sleep(500 * time.Millisecond)
		ch5 <- "来自channel2"
	}()
	
	// 使用select监听多个channel
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch4:
			fmt.Println("接收到:", msg1)
		case msg2 := <-ch5:
			fmt.Println("接收到:", msg2)
		case <-time.After(2 * time.Second):
			fmt.Println("超时")
		}
	}
}

// 生产者函数
func producer(ch chan int) {
	for i := 1; i <= 5; i++ {
		ch <- i
		fmt.Printf("生产: %d\n", i)
		time.Sleep(100 * time.Millisecond)
	}
	close(ch) // 关闭channel
}

// 消费者函数
func consumer(ch chan int) {
	for value := range ch { // 使用range接收直到channel关闭
		fmt.Printf("消费: %d\n", value)
		time.Sleep(200 * time.Millisecond)
	}
} 