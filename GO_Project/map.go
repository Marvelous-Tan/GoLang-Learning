package main

import "fmt"

func main() {
	map1 := make(map[string]int)
	// map无序，key-value对
	// map的key不能重复，value可以重复
	map1["a"] = 1
	map1["b"] = 2
	fmt.Println(map1)

	// map的添加和删除
	map1["c"] = 3
	delete(map1, "a")
	fmt.Println(map1)

	// map的遍历
	for key, value := range map1 {
		fmt.Println(key, value)
	}

	// map的查找
	value, ok := map1["b"]
	if ok {
		fmt.Println(value)
	}
	
	// map的删除
	delete(map1, "b")
	fmt.Println(map1)

	// map的清空
	map1 = make(map[string]int)
	fmt.Println(map1)
	
	// map修改
	map1["a"] = 100
	fmt.Println(map1)

}