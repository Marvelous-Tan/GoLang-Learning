package main

import "fmt"

const (
	// 常量值具有只读属性，不能被修改
	pi = 3.14

	// iota 是go语言的常量计数器，从0开始，每行自增1
	// 第一行的iota的值为0
	BEIJING = iota
	SHANGHAI
	GUANGDONG
	

)

func main() {
	fmt.Println(pi)
	fmt.Println(BEIJING)
	fmt.Println(SHANGHAI)
	fmt.Println(GUANGDONG)
	fmt.Println("BEIJING = ", BEIJING)
	fmt.Println("SHANGHAI = ", SHANGHAI)
	fmt.Println("GUANGDONG = ", GUANGDONG)


	
	
}