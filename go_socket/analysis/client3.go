package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", ":5000")
	if err != nil {
		fmt.Println(err)
		return
	}

	data := os.Args[1]

	fmt.Printf("try to write with TCP: %s\n", data)

	conn.Write([]byte(data))

	time.Sleep(10 * time.Second)
}
