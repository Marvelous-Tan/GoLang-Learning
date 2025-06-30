package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string `info:"name" doc:"姓名"`
	Sex string `info:"sex" doc:"性别"`
}

// Tag 是结构体中的字段对应的标签
// Get 是获取结构体中的字段对应的标签
// NumField 是获取结构体中的字段数量
// Field 是获取结构体中的字段
func findTag(str interface{}) {
	t := reflect.TypeOf(str).Elem()
	for i := 0; i < t.NumField(); i++ {
		taginfo := t.Field(i).Tag.Get("info")
		tagdoc := t.Field(i).Tag.Get("doc")
		fmt.Println("info :", taginfo, "doc :", tagdoc)
	}
}

func main() {
	findTag(Person{})
}	