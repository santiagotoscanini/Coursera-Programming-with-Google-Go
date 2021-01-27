package main

import (
	"fmt"
	"sort"
	"strconv"
)

/*
Write a program which prompts the user to enter integers and stores the integers in a sorted slice.
The program should be written as a loop. Before entering the loop, the program should create an empty integer slice of size (length) 3.
During each pass through the loop, the program prompts the user to enter an integer to be added to the slice.
The program adds the integer to the slice, sorts the slice, and prints the contents of the slice in sorted order.
The slice must grow in size to accommodate any number of integers which the user decides to enter.
The program should only quit (exiting the loop) when the user enters the character ‘X’ instead of an integer.
*/

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
