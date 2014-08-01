package main

import "fmt"

func main() {
	i := 7
	switch i {
	case 1:
		fmt.Println(i)
	case 2, 3:
		fmt.Println(i) //匹配多种参数
	case 4:
		fmt.Println(i)
		fallthrough
		//go里case后面自动break，加上fallthrough，强制继续执行
	case 5:
		fmt.Println(i)
	default:
		fmt.Println("default case")
	}
}
