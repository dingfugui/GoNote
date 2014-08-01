// testMap
package main

import (
	"fmt"
)

func main() {
	var maps map[string]int     //声明1
	maps = make(map[string]int) //声明2

	maps["one"] = 1 //赋值
	maps["ten"] = 10

	fmt.Println(maps)
	fmt.Println(maps["one"])
	fmt.Println(maps["two"])

	dict := map[string]int{"C": 4, "java": 5}
	value, f := dict["C#"]
	fmt.Println(f)
	fmt.Println(value)
}
