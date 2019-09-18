package main

import (
	"fmt"
	"net"
)

func main() {
	l, err := net.Listen("tcp", ":5000")
	if err != nil {
		fmt.Println(err)
	}

	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			break
		}

		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {

	fmt.Println("accept a new connect")

	for {
		var buf = make([]byte, 100)

		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println(err)
			break
		}

		if n > 0 {
			fmt.Printf("accept buf: %s\n", buf)
		}
	}

	defer conn.Close()
}
