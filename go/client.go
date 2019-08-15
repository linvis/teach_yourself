package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

func main() {
	if len(os.Args[1:]) < 1 {
		fmt.Println("invalid addr")
		os.Exit(1)
	}

	service := os.Args[1]

	tcpAddr, _ := net.ResolveTCPAddr("tcp4", service)

	conn, _ := net.DialTCP("tcp", nil, tcpAddr)

	conn.Write([]byte("hello"))

	result, _ := ioutil.ReadAll(conn)

	fmt.Println(string(result))
}
