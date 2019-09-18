package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	l, err := net.Listen("tcp", ":5000")
	if err != nil {
		fmt.Println(err)
	}

	i := 0
	for {
		time.Sleep(10 * time.Second)
		_, err := l.Accept()
		if err != nil {
			fmt.Println(err)
		}

		i++
		fmt.Printf("server accept %d\n", i)
	}
}
