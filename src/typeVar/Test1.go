// Test1 project main.go
package main

import (
	"fmt"
)

//变量
func main() {
	fmt.Println("Hello蛤蛤!") //go 使用UTF-8编码，支持多语言
	var a int               //定义一个int类型的变量
	var a1, a2, a3 int      //定义多个变量
	var b int = 1
	var c1, c2 = 2, 3
	var d = "abc"  //自动识别类型
	e := d         //无需 var 自动推断
	_, b := 34, 35 //_是个特殊变量名，任何赋予它的值都会被丢弃

}
