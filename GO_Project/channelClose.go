package main

import "fmt"

func main() {
	ch := make(chan int, 3)
	for i := 0; i < 3; i++ {
		ch <- i
	}
    // close 一个channel后，不能再向channel中写入数据，但是可以继续从channel中读取数据
	close(ch)

	for i := 0; i < 4; i++ {
		data, ok := <-ch
		// ok 为 false 表示channel已经关闭
		// ok 为 true 表示channel还没有关闭
		if ok {
			fmt.Println("data = ", data)
		} else {
			fmt.Println("channel closed")
			break
		}
	}
}