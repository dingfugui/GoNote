package main

import "fmt"

var arr [10]int //定义数组

func main() {
	arr[0] = 1
	fmt.Println(arr[9])

	a := [3]int{1, 2, 3}
	b := [10]int{1, 2, 3}
	c := [...]int{3, 5, 4} //自动根据个数设置长度
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
