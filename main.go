package main

import (
	"fmt"
	"os"
	"strings"
)

func WordCount(str string) int {
	length := len(strings.Fields(str))
	return length
}

func main() {
	str := os.Args[1]
	length := WordCount(str)

	fmt.Printf("word count in string '%v' is %v.\n", str, length)
}
