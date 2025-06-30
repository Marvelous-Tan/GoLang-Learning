package main

import (
	"fmt"
	"reflect"
)

func reflectNum(arg interface{}) {
	v := reflect.ValueOf(arg)
	fmt.Println("type :", v.Type())
	fmt.Println("value :", v.Interface())
	// v.Interface() 返回的是一个空接口，可以接收任何类型的数据，所以需要类型断言
	fmt.Println("value :", reflect.ValueOf(arg))
	fmt.Println("type :", reflect.TypeOf(arg))
}

func main() {
	num := 1.2345
	reflectNum(num)
}