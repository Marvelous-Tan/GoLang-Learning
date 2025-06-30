package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5}
	// 动态数组，长度不固定，可以动态添加元素
	// 切片是数组的一部分，切片是数组的引用
	// 切片是动态数组，长度不固定，可以动态添加元素
	// 切片是数组的一部分，切片是数组的引用
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}
	for index, value := range slice {
		fmt.Println(index, value)
	}

	fmt.Printf("len(slice): %d\n", len(slice))
	fmt.Printf("cap(slice): %d\n", cap(slice))

	slice = append(slice, 6)
	fmt.Printf("len(slice): %d\n", len(slice))
	fmt.Printf("cap(slice): %d\n", cap(slice))
	// cap作用：cap是容量，表示切片底层数组的长度
	
	// slice的追加和截取
	slice = append(slice, 7, 8, 9)
	fmt.Println(slice)

	slice2 := slice[0:3]// 3是切片的长度，0是切片的起始位置，3是切片的结束位置
	fmt.Println(slice2)
	
}






//Println 和 Printf 的区别
//Println 会自动换行
//Printf 不会自动换行
//Println 会自动添加空格
//Printf 不会自动添加空格
//Println 会自动添加换行符
//Printf 不会自动添加换行符
//Println 会自动添加换行符