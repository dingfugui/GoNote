package main

import (
	"fmt"
)

//数值类型
/*
	整数类型有无符号和带符号两种。
	Go同时支持int和uint，这两种类型的长度相同，但具体长度取决于不同编译器的实现。
	Go里面也有直接定义好位数的类型：
	rune, int8, int16, int32, int64
	和
	byte, uint8, uint16, uint32, uint64。
	其中rune是int32的别称，byte是uint8的别称
*/
/*
	浮点数类型有float32 和 float64
	没有float类型
	默认是float64
*/
func main() {
	fmt.Println("Hello World!")
	var a int8
	var b int16
	c := a + b //不允许不同类型变量之间赋值或操作，即，没有java的隐式转换
}
