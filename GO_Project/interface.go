package main

import "fmt"

// 通用万能接口，可以接收任何类型的数据
// 可以实现类型断言

// 本质是个指针
// 可以实现多态
type Animal interface {
	Sleep()
	GetColor() string
	GetType() string
}

type Cat struct {	
	color string
}

func (this *Cat) Sleep() {
	fmt.Println("Cat is sleeping")
}

func (this *Cat) GetColor() string {
	return this.color
}		

func (this *Cat) GetType() string {
	return "Cat"
}

func myfunc(arg interface{}) {
	fmt.Println("myfunc")
	fmt.Println(arg)
}

func main() {
	animal := Cat{color: "white"}
	animal.Sleep()
	fmt.Println(animal.GetColor())
	fmt.Println(animal.GetType())

	value, ok := animal.(Cat)
	if !ok {
		fmt.Println("convert failed")
		return
	}
	fmt.Println(value)
}