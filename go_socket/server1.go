package main

import (
	"fmt"
	"net"
)

func handleConn(conn net.Conn) {
	fmt.Println("accept a new connect")

	defer conn.Close()
}

func main() {
	l, err := net.Listen("tcp", ":5000")
	if err != nil {
		fmt.Println(err)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println(err)
		}

		go handleConn(conn)
	}
}
