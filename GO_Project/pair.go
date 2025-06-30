package main

import "fmt"

func main() {
	var a string
	a = "hello"
	var alltype interface{}
	alltype = a
	str, ok := alltype.(string)
	if !ok {
		fmt.Println("convert failed")
		return
	}
	fmt.Println(str)

	// 类型断言
}
