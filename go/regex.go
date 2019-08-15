package main

import (
	"fmt"
	"regexp"
)

func main() {
	a := "I am learning Go language"

	re, _ := regexp.Compile("[a-z]{2,4}")

	one := re.Find([]byte(a))
	fmt.Println("Find: ", string(one))

	two := re.FindAll([]byte(a), -1)
	fmt.Println("Find: ", string(two[0]))
}
