package main

import "fmt"

func main() {
	ch := make(chan int)

	go func() {

		defer fmt.Println("goroutine done")
		fmt.Println("goroutine doing")
		ch <- 666
	}()

	data := <-ch
	fmt.Println("data = ", data)
	fmt.Println("main done")
}