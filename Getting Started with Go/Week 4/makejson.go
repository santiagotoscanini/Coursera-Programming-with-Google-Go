package main

import (
	"encoding/json"
	"fmt"
)

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
