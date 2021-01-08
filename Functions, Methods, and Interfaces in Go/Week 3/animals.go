package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (a *Animal) Eat() {
	fmt.Println(a.food)
}

func (a *Animal) Move() {
	fmt.Println(a.locomotion)
}

func (a *Animal) Speak() {
	fmt.Println(a.noise)
}

var cow = Animal{
	food:       "grass",
	locomotion: "walk",
	noise:      "moo",
}
var bird = Animal{
	food:       "worms",
	locomotion: "fly",
	noise:      "peep",
}
var snake = Animal{
	food:       "mice",
	locomotion: "slither",
	noise:      "hsss",
}

func main() {
	for {
		fmt.Print("> ")
		inputScanner := bufio.NewScanner(os.Stdin)
		inputScanner.Scan()

		words := strings.Fields(inputScanner.Text())

		var animalSelected Animal
		switch words[0] {
		case "cow":
			animalSelected = cow
		case "bird":
			animalSelected = bird
		case "snake":
			animalSelected = snake
		}

		switch words[1] {
		case "eat":
			animalSelected.Eat()
		case "move":
			animalSelected.Move()
		case "speak":
			animalSelected.Speak()
		}
	}
}
