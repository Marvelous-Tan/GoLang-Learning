package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 3)
	

	fmt.Println("len(ch) = ", len(ch), "cap(ch) = ", cap(ch))
	// len 表示当前channel中剩余未读取的数据个数
    // cap 表示channel的容量

	go func() {
		defer fmt.Println("goroutine done")
		for i := 0; i < 4; i++ {
			ch <- i
			fmt.Println("goroutine sending: ", i, "len(ch) = ", len(ch), "cap(ch) = ", cap(ch))
		}
	}()

	time.Sleep(1 * time.Second)
	
	for i := 0; i < 4; i++ {
		data := <-ch
		fmt.Println("data = ", data)
	}
	time.Sleep(1 * time.Second)
	fmt.Println("main done")
}