package main

import (
	"fmt"
	"reflect"
)

// 建议不加分号
func main() {
	var a int = 10
	

	fmt.Println("type of a is", reflect.TypeOf(a))
    var abc = "abc"
	fmt.Println("type of abc is", reflect.TypeOf(abc))

	e :=100
	g := 3.14
	fmt.Println("type of e is", reflect.TypeOf(e))
	fmt.Println(e)

	fmt.Println("type of g is", reflect.TypeOf(g))
	fmt.Printf("type of g is = %f\n", g)
	


	
}

