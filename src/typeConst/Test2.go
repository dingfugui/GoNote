package main

import (
	"fmt"
)

//常量
func main() {
	fmt.Println("Hello World!")
	//常量在编译阶段就确定了，运行时无法改变，可以是数值、布尔、字符串
	const c1 = "a"
	const Pi float32 = 3.1415
}
