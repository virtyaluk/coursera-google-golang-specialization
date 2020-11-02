package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Please, enter a string")

	ioReader := bufio.NewReader(os.Stdin)
	inputStr, err := ioReader.ReadString('\n')

	if err != nil {
		fmt.Println("Input error")
	}

	inputStr = strings.TrimSuffix(inputStr, "\n")
	inputStr = strings.ToLower(inputStr)

	hasI := strings.HasPrefix(inputStr, "i")
	hasA := strings.Contains(inputStr, "a")
	hasN := strings.HasSuffix(inputStr, "n")

	if hasI && hasA && hasN {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
