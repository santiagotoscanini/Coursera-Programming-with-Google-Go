package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {

	var s string
	var sli = make([]int, 0, 3)
	for {
		fmt.Printf("Enter an integer: ")
		fmt.Scan(&s)
		x, err := strconv.Atoi(s)
		if err == nil { /* if the input from user is a number */
			sli = append(sli, x)
			sort.Ints(sli)
			fmt.Println(sli)
		} else if s == "X" { /* if the user enters 'X' */
			break
		} else { /* if the user enters anything else other than a number or 'X', print the same slice */
			fmt.Println(sli)
		}
	}
}
