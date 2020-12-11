package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := strings.ToUpper(scanner.Text())

	if strings.IndexRune(input, 'I') == 0 &&
		strings.LastIndexByte(input, 'N') == len(input)-1 &&
		strings.ContainsRune(input, 'A') {
		fmt.Print("Found!")
	} else {
		fmt.Print("Not Found!")
	}
}
