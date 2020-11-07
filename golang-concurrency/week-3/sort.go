package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

func takeInput() []string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter comma separated list of integers you want to sort:")

	scanner.Scan()

	return strings.Fields(scanner.Text())
}

func subSort(wg *sync.WaitGroup, nums []int) {
	fmt.Println(nums)
	sort.Ints(nums)

	if wg != nil {
		wg.Done()
	}
}

func merge(to, from1, from2 []int) {
	i, j, k := 0, 0, 0

	for i < len(from1) && j < len(from2) {
		if from1[i] < from2[j] {
			to[k] = from1[i]
			i++
		} else {
			to[k] = from2[j]
			j++
		}

		k++
	}

	for i < len(from1) {
		to[k] = from1[i]
		i++
		k++
	}

	for j < len(from2) {
		to[k] = from2[j]
		j++
		k++
	}
}

func main() {
	nums := []int{}

	for _, token := range takeInput() {
		num, _ := strconv.Atoi(token)
		nums = append(nums, num)
	}

	if len(nums) < 4 {
		subSort(nil, nums)
		fmt.Println("Sorted array is:", nums)

		return
	}

	var wg sync.WaitGroup
	chunkSize := len(nums) / 4
	chunk1, chunk2, chunk3, chunk4 := nums[:chunkSize], nums[chunkSize:2 * chunkSize], nums[2 * chunkSize:3 * chunkSize], nums[3 * chunkSize:]

	wg.Add(4)
	go subSort(&wg, chunk1)
	go subSort(&wg, chunk2)
	go subSort(&wg, chunk3)
	go subSort(&wg, chunk4)
	wg.Wait()

	chunk12, chunk34, res := make([]int, len(chunk1) + len(chunk2)), make([]int, len(chunk3) + len(chunk4)), make([]int, len(nums))
	merge(chunk12, chunk1, chunk2)
	merge(chunk34, chunk3, chunk4)
	merge(res, chunk12, chunk34)

	fmt.Println("Sorted array is:", res)
}