package main

import (
	_ "GO_Project/lib/lib1"
	mylib2 "GO_Project/lib/lib2"
	// ."mylib2 "GO_Project/lib/lib2"
	// 添加了.之后，在调用Init2()时，可以不使用mylib2.Init2()，直接使用Init2()
	// 尽量不要用.
	// 可能lib1和lib2有同名函数产生冲突
)

func main() {
	//lib1.Init1()
	mylib2.Init2()
}