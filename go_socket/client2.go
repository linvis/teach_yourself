package main

import (
	"fmt"
	"net"
)

func main() {
	for i := 0; i < 1000; i++ {
		_, err := net.Dial("tcp", ":5000")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("%d create a new connect\n", i)
	}
}
