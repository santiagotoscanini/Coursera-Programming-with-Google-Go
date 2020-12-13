package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	maxLength = 20
)

type Name struct {
	firstName string
	lastName  string
}

func (n *Name) Set(firstName string, lastName string) {
	var rs []rune

	n.firstName = firstName
	if len(firstName) > maxLength {
		rs = []rune(firstName)
		n.firstName = string(rs[:maxLength])
	}

	n.lastName = lastName
	if len(lastName) > maxLength {
		rs = []rune(lastName)
		n.lastName = string(rs[:maxLength])
	}
}

func main() {
	inputScanner := bufio.NewScanner(os.Stdin)
	inputScanner.Scan()
	// File path are relative from root of repo, e.g: "Getting Started with Go/Week 4/test.txt"
	fileName := strings.ToUpper(inputScanner.Text())

	f, _ := os.Open(fileName)
	defer f.Close()
	scanner := bufio.NewScanner(f)

	var names []Name

	for scanner.Scan() {
		name := strings.Fields(scanner.Text())
		nameStruct := &Name{}
		nameStruct.Set(name[0], name[1])

		names = append(names, *nameStruct)
	}

	for _, name := range names {
		fmt.Printf("%s, %s\n", name.firstName, name.lastName)
	}
}
