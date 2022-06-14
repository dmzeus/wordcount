package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	inputString := os.Args[1]
	length := len(strings.Fields(inputString))

	fmt.Println(length)
}
