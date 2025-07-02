package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("=== Goroutine 演示 ===")
	
	// 1. 基本的 goroutine 创建
	fmt.Println("1. 基本 goroutine:")
	go func() {
		fmt.Println("这是一个协程")
	}()
	
	// 2. 带参数的 goroutine
	fmt.Println("2. 带参数的 goroutine:")
	for i := 1; i <= 3; i++ {
		go func(id int) {
			fmt.Printf("协程 %d 正在运行\n", id)
		}(i)
	}
	
	// 3. 调用普通函数作为 goroutine
	fmt.Println("3. 调用函数作为 goroutine:")
	go printMessage("Hello from goroutine!")
	
	// 4. 等待一段时间让 goroutine 执行
	fmt.Println("4. 等待 goroutine 执行...")
	time.Sleep(1 * time.Second)
	
	// 5. 演示并发执行
	fmt.Println("5. 并发执行演示:")
	start := time.Now()
	
	// 启动多个 goroutine
	for i := 1; i <= 5; i++ {
		go func(id int) {
			time.Sleep(100 * time.Millisecond)
			fmt.Printf("任务 %d 完成\n", id)
		}(i)
	}
	
	time.Sleep(200 * time.Millisecond)
	fmt.Printf("所有任务完成，耗时: %v\n", time.Since(start))
}

// 普通函数
func printMessage(msg string) {
	fmt.Println(msg)
} 