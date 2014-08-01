package main

import (
	"errors"
	"fmt"
)

func main() {
	err := errors.New("异常信息")
	if err != nil {
		fmt.Println(err)
	}
}
