package main

import "fmt"
import "net"

func main() {
	name := "192.168.104.1"
	addr := net.ParseIP(name)

	if addr == nil {
		fmt.Println("Invalid adress")
	} else {
		fmt.Println("ip is", addr.String())
	}

}
