package main

import "fmt"

// 自定义一个类型
type myInt int

// 定义一个结构体包含变量和方法
// 类的首字母大写，表示公有，小写表示私有
type Person struct {
	Name string
	Age int
	Sex string
}

func (this *Person) getAge() int {
	return this.Age
}
func (this *Person) getName() string {
	return this.Name
}
func (this Person) getSex() string {
	return this.Sex
}





// 包含引用传递的函数
func print(p Person) {
	fmt.Println(p)
}
// 包含值传递的函数
func (p Person) print2() {
	fmt.Println(p)
}


func main() {
	person := Person{Name: "张三", Age: 20, Sex: "男"}
	fmt.Println(person)
	fmt.Println(person.getAge())
	fmt.Println(person.getName())
	fmt.Println(person.getSex())

	print(person)
	person.print2()
}