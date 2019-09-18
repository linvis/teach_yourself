package main

import (
	"fmt"
	"net"
)

func main() {
	_, err := net.Dial("tcp", ":5000")
	if err != nil {
		fmt.Println(err)
		return
	}
}
