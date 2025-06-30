package main

import "fmt"

func main() {
	// defer fmt.Println("main end1")
	// defer fmt.Println("main end2")
	// fmt.Println("main start")
	// defer 是后进先出
	// 先执行main end2，再执行main end1
	// 先执行main end2，再执行main end1	

	fmt.Println(test())
	
}

// return 和 defer 的执行顺序
func test() int {
	defer fmt.Println("defer")
	return 1
}
