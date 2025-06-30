package main

import "fmt"

type Person struct {
	Name string
	Age int
	Sex string
}

type Student struct {
	Person
	StudentID string
}


func (this *Person) GetName() string {
	return this.Name
}

// 重定义方法
func (this *Student) GetName() string {
	return this.Name + "学生"
}

func main() {
	student := Student{Person: Person{Name: "张三", Age: 20, Sex: "男"}, StudentID: "123456"}
	fmt.Println(student.GetName())
}