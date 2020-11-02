package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var inputStr string
	slice, keepLoop := make([]int, 3), true

	for keepLoop {
		fmt.Println("Please, enter an integer to be added to the slice:")

		_, err := fmt.Scan(&inputStr)

		if err != nil {
			fmt.Println("Invalid imput")
		}

		if strings.HasPrefix(strings.ToLower(inputStr), "x") {
			keepLoop = false
			fmt.Println("Exiting...")
			continue
		}

		num, err := strconv.Atoi(inputStr)

		if err != nil {
			fmt.Println("Unable to parse given integer")
			continue
		}

		slice = append(slice, num)
		
		sort.Ints(slice)
		fmt.Println("Current slice: ", slice)
	}
}