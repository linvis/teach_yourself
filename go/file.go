package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	arg := os.Args[1]
	data, _ := ioutil.ReadFile(arg)

	fmt.Printf(string(data))
}
