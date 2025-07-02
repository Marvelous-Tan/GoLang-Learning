package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("goroutine sending: ", i)
			ch <- i
		}
		close(ch)
	}()

	// range 会自动关闭channel
	// 如果channel没有关闭，range会一直等待，直到channel关闭
	// 如果channel已经关闭，range会自动退出
	for data := range ch {
		fmt.Println("data = ", data)
		time.Sleep(time.Second)
	}

	fmt.Println("main done")

}