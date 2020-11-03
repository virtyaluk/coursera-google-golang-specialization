package main

import (
	"bufio"
	"fmt"
	"strings"
	"os"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {}

func (c Cow) Eat() {
	fmt.Println("grass")
}

func (c Cow) Move() {
	fmt.Println("walk")
}

func (c Cow) Speak() {
	fmt.Println("moo")
}

type Bird struct {}

func (b Bird) Eat() {
	fmt.Println("worms")
}

func (b Bird) Move() {
	fmt.Println("fly")
}

func (b Bird) Speak() {
	fmt.Println("peep")
}

type Snake struct {}

func (s Snake) Eat() {
	fmt.Println("mice")
}

func (s Snake) Move() {
	fmt.Println("slither")
}

func (s Snake) Speak() {
	fmt.Println("hsss")
}

func takeInput() []string {
	scanner := bufio.NewScanner(os.Stdin)
	
	fmt.Printf("> ")
	scanner.Scan()

	return strings.Fields(scanner.Text())
}

func main() {
	store := map[string]string{}

	for true {
		var cmd, name, info string
		args, invalid := takeInput(), false

		if len(args) == 3 {
			cmd, name, info = strings.ToLower(args[0]), strings.ToLower(args[1]), strings.ToLower(args[2])
		} else {
			invalid = true
		}

		if cmd == "newanimal" {
			store[name] = info
			fmt.Println("Created it!")
		} else if cmd == "query" {
			var a Animal

			if store[name] == "cow" {
				a = Cow{}
			} else if store[name] == "bird" {
				a = Bird{}
			} else if store[name] == "snake" {
				a = Snake{}
			} else {
				invalid = true
			}

			if info == "eat" {
				a.Eat()
			} else if info == "move" {
				a.Move()
			} else if info == "speak" {
				a.Speak()
			} else {
				invalid = true
			}
		}

		if invalid {
			fmt.Println("Invalid input")
		}
	}
}
