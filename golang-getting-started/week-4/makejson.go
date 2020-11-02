package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	person := make(map[string]string)

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Please, enter a name:")
	scanner.Scan()

	person["name"] = scanner.Text()

	fmt.Println("Please, enter an address:")
	scanner.Scan()

	person["address"] = scanner.Text()

	jsonData, err := json.Marshal(person)

	if err != nil {
		fmt.Println("Unable to convert to JSON. Exiting...")
		return
	}

	fmt.Println("Resulting JSON: ", string(jsonData))
}
