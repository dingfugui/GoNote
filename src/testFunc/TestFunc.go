package main

import "fmt"

func main() {
	//fmt.Println(foo1(1, 2))
	foo2(1, 23, 4)
}

//多返回值
func foo1(a, b int) (int, int) {
	return b, a
}

//可变参
func foo2(arg ...int) {
	for _, n := range arg {
		fmt.Println(n)
	}
}
