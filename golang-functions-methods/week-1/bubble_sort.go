package main

import (
	"fmt"
	"strconv"
)

func main() {
	ar, token := []int{}, ""

	fmt.Println("You are going to be prompted to enter up to 10 integers.")

	for i := 0; i < 10; i++ {
		fmt.Println("Please, enter an integer (or 'x' to proceed with sorting):")
		_, err := fmt.Scan(&token)

		if err != nil {
			fmt.Println("Invalid input. Try again.")
		} else if token == "x" {
			break
		} else if num, err := strconv.Atoi(token); err == nil {
			ar = append(ar, num)
		}
	}

	BubbleSort(ar)

	fmt.Println("Resulting array sorted using bubble sort:", ar)
}

func BubbleSort(ar []int) {
	sorted := false

	for !sorted {
		swapped := false

		for i := 0; i < len(ar) - 1; i++ {
			if ar[i] > ar[i + 1] {
				Swap(ar, i)
				swapped = true
			}
		}

		sorted = !swapped
	}
}

func Swap(ar []int, i int) {
	ar[i], ar[i + 1] = ar[i + 1], ar[i]
}
