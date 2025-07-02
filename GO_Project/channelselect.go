package main

import "fmt"

func main() {
	ch := make(chan int)
	ch2 := make(chan int)

	go func() {
		ch <- 1
	}()

	select {
	case <-ch:
		fmt.Println("ch")
	case <-ch2:
		fmt.Println("ch2")
	}
}