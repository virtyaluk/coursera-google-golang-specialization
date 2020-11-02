package main

import (
	"bufio"
	"fmt"
	"strings"
	"os"
)

type Person struct {
	fname, lname string
}

func main() {
	names := []Person{}
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Please, enter the name of the text file:")
	scanner.Scan()

	fName := scanner.Text()

	fd, err := os.Open(fName)

	if err != nil {
		fmt.Println("Error while oppening the file")
		return
	}

	scanner = bufio.NewScanner(fd)

	for scanner.Scan() {
		personaName := strings.Fields(scanner.Text())
		names = append(names, Person{personaName[0], personaName[1]})
	}

	fmt.Println("The names from the file are:")

	for _, name := range names {
		fmt.Println(name.fname, name.lname)
	}
}
