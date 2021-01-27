package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
Write a program which reads information from a file and represents it in a slice of structs.
Assume that there is a text file which contains a series of names.
Each line of the text file has a first name and a last name, in that order, separated by a single space on the line.

Your program will define a name struct which has two fields, fname for the first name, and lname for the last name.
Each field will be a string of size 20 (characters).

Your program should prompt the user for the name of the text file.
Your program will successively read each line of the text file and create a struct which contains the first and last names found in the file.
Each struct created will be added to a slice, and after all lines have been read from the file, your program will have a slice containing one struct for each line in the file.
After reading all lines from the file, your program should iterate through your slice of structs and print the first and last names found in each struct.
*/

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
	// File path are relative from root of repo, e.g: "1. Getting Started with Go/Week 4/test.txt"
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
