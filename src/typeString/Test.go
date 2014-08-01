package main

import (
	"fmt"
)

var s = "abc"
var s1 string
var s2 string = ""

//变量
func main() {
	fmt.Println("Hello蛤蛤!")
	fmt.Println(s[0])
	//s[0] = 'a'//字符串是不允许修改的
	c := []byte(s)
	c[0] = 'w'
	ss := string(c)
	fmt.Println(ss) //如果非要修改，只能先转换成stirng

	s = s1 + s2 //用+拼接

	s = "c" + s[1:] //字符串不能修改，但是可以切片
	fmt.Println(s)
	
	s3 :='a
	b'//声明多行字符串
}

func test() {
	a1 := "a1"
	s = "a2"
	fmt.Println(a1)
}
