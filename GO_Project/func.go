package main

import (
	"fmt"
	"reflect"
)

func test() float64 {
    var(
	    a =100.0
	    b=3.14
    ) 
 return a+b
}

func test2(a int,b int) (int,float64) {
	return a+b,float64(a+b)
}

func main() {
	fmt.Println(test())
	a :=10
	b :=20
	c,d :=test2(a,b)
	fmt.Println(c,d)
	fmt.Println("type of c is",reflect.TypeOf(c))
	fmt.Println("type of d is",reflect.TypeOf(d))
}