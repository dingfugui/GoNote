package main

import "fmt"
import "net"
import "os"
import "io/ioutil"

func main() {
	//name := "www.baidu.com:80"
	name := "192.168.106.229:7777"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", name)
	checkError(err)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)
	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkError(err)
	result, err := ioutil.ReadAll(conn)
	checkError(err)
	fmt.Println(string(result))
}
func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error:%s", err.Error())
		os.Exit(1)
	}
}
