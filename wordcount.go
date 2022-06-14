package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	inputString := os.Args[1]
	length := len(strings.Fields(inputString))

	fmt.Printf("word count in string '%v' is %v.\n", inputString, length)
}
