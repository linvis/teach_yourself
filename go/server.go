package main

import (
	"net"
	"time"
)

func main() {
	service := ":7777"
	tcpAddr, _ := net.ResolveTCPAddr("tcp4", service)

	listener, _ := net.ListenTCP("tcp", tcpAddr)

	for {
		conn, _ := listener.Accept()
		go handleTime(conn)
	}
}

func handleTime(conn net.Conn) {
	daytime := time.Now().String()
	conn.Write([]byte(daytime))
	conn.Close()
}
