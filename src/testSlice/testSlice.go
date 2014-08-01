// testSlice
package main

import (
	"fmt"
)

//slice 动态数组
func main() {
	s := []byte{'a', 'b', 'c', 'd'}

	var a, b []byte //slice即不指定数组长度

	a = s[2:3]
	b = s[1:3]

	fmt.Println(a)
	fmt.Println(b)

	fmt.Println(len(s))
	fmt.Println(append(s, 1, 2))
	fmt.Println(cap(s))
	var c []byte
	fmt.Println(copy(c, s))
}
