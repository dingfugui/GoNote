package main

import "fmt"

var i = 1

func main() {
	//foo()
	testFor()
	if i > 1 {
		fmt.Println(i)
	}

	if x := 3; x > 1 {
		fmt.Println(i)
	}

}
func foo() {
	i := 0
here:
	i++
	if i < 3 {
		fmt.Println("i:", i)
		goto here //goto跳转到标签
	}
}

func testFor() {
	for i := 0; i < 10; i++ {
		if i == 1 {
			continue
		}
		if i == 2 {
			break
		}
	}

	dict := map[int]string{1: "a", 2: "b"}
	//dict := map[string]int{"C": 4, "java": 5}

	for k, v := range dict {
		fmt.Print(k, "~")
		fmt.Println(v)
	}
}
