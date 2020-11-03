package main

import (
	"bufio"
	"fmt"
	"strings"
	"os"
)

type Animal struct {
	food, locomotion, noise string
}

func (a Animal) Eat() {
	fmt.Println(a.food)
}

func (a Animal) Move() {
	fmt.Println(a.locomotion)
}

func (a Animal) Speak() {
	fmt.Println(a.noise)
}

func takeInput() (string, string) {
	scanner := bufio.NewScanner(os.Stdin)
	
	fmt.Printf(">")
	scanner.Scan()

	fields := strings.Fields(scanner.Text())

	return strings.ToLower(fields[0]), strings.ToLower(fields[1])
}

func main() {
	animals := map[string]Animal{}
	animals["cow"] = Animal{"grass", "walk", "moo"}
	animals["bird"] = Animal{"worms", "fly", "peep"}
	animals["snake"] = Animal{"mice", "slither", "hsss"}

	for true {
		animal, info := takeInput()
		a, ok := animals[animal]

		ok = ok && (info == "eat" || info == "move" || info == "speak")

		if !ok {
			fmt.Println("Invalid input.")
			continue
		}

		if info == "eat" {
			a.Eat()
		} else if info == "move" {
			a.Move()
		} else {
			a.Speak()
		}
	}
}
