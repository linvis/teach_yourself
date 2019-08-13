package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	tick := time.Tick(1 * time.Second)
	abort := make(chan struct{})

	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()

	for count := 10; count > 0; count-- {
		select {
		case <-abort:
			fmt.Println("abort")
			return
		case <-time.After(10 * time.Second):
			fmt.Println("done")
		case <-tick:
			fmt.Println("cur tick ", count)
		}
	}
}
