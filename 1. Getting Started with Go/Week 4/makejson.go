package main

import (
	"encoding/json"
	"fmt"
)

/*
Write a program which prompts the user to first enter a name, and then enter an address.
Your program should create a map and add the name and address to the map using the keys “name” and “address”, respectively.
Your program should use Marshal() to create a JSON object from the map, and then your program should print the JSON object.
*/

func main() {
	var name string
	_, _ = fmt.Scan(&name)

	var address string
	_, _ = fmt.Scan(&address)

	person := map[string]string{
		"name":    name,
		"address": address,
	}

	var barr []byte
	barr, _ = json.Marshal(person)
	fmt.Printf("%s", barr)
}
