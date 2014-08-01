package main

import "fmt"

/*
	Go 的默认原则
	字母或方法
	大写字母开头的相当于public
	小写字母开头的相当于private
*/

//导包、声明，可以单独一行，也可以分组
var (
	i  int
	pi float32
)

const (
	x = iota //0
	y = iota //1
	z = iota //2
	w        //3
	//常量声明时，不写类型，默认和之前一个值的字面相同
	//即，此处w为iota，所以w=3
)
const v = iota //每遇到一个关键字，iota重置
const (
	e, f, g = iota, iota, iota //iota如果在一行，不会自增
)

func main() {
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println(w)
	fmt.Println(v)
	fmt.Println(e)
	fmt.Println(f)
}
