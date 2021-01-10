package main

import (
	"fmt"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct{}

func (Cow) Eat() {
	fmt.Println("grass")
}

func (Cow) Move() {
	fmt.Println("walk")
}

func (Cow) Speak() {
	fmt.Println("moo")
}

type Bird struct{}

func (Bird) Eat() {
	fmt.Println("worms")
}

func (Bird) Move() {
	fmt.Println("fly")
}

func (Bird) Speak() {
	fmt.Println("peep")
}

type Snake struct{}

func (Snake) Eat() {
	fmt.Println("mice")
}

func (Snake) Move() {
	fmt.Println("slither")
}

func (Snake) Speak() {
	fmt.Println("hsss")
}

var animals = make(map[string]Animal, 10)

func main() {
	for {
		var command, animalName, request string
		fmt.Print("> ")
		fmt.Scan(&command, &animalName, &request)

		if command == "newanimal" {
			switch request {
			case "cow":
				animals[animalName] = Cow{}
			case "bird":
				animals[animalName] = Bird{}
			case "snake":
				animals[animalName] = Snake{}
			default:
				fmt.Println("Incorrect animal type!")
			}
			fmt.Println("Created it!")
		} else if command == "query" {
			ani, exists := animals[animalName]
			if !exists {
				fmt.Printf("There is no animal named %s", animalName)
				fmt.Println()
			} else {
				switch request {
				case "eat":
					ani.Eat()
				case "move":
					ani.Move()
				case "speak":
					ani.Speak()
				default:
					fmt.Println("Incorrect animal information")
				}
			}
		} else {
			fmt.Println("Incorrect command!")
		}
	}
}
