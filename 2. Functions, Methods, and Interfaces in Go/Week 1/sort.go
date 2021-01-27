package main

import (
	"fmt"
	"strconv"
)

/*
Write a Bubble Sort program in Go. The program should prompt the user to type in a sequence of up to 10 integers.
The program should print the integers out on one line, in sorted order, from least to greatest.
Use your favorite search tool to find a description of how the bubble sort algorithm works.

As part of this program, you should write a function called BubbleSort() which takes a slice of integers as an argument and returns nothing.
The BubbleSort() function should modify the slice so that the elements are in sorted order.

A recurring operation in the bubble sort algorithm is the Swap operation which swaps the position of two adjacent elements in the slice.
You should write a Swap() function which performs this operation.
Your Swap() function should take two arguments, a slice of integers and an index value i which indicates a position in the slice.
The Swap() function should return nothing, but it should swap the contents of the slice in position i with the contents in position i+1.
*/

func main() {
	sli := ReadSliceFromInput()
	BubbleSort(sli)

	fmt.Println(sli)
}

func ReadSliceFromInput() []int {
	var s string
	var sli = make([]int, 0, 10)

	for i := 0; i < 10; i++ {
		_, _ = fmt.Scan(&s)
		x, err := strconv.Atoi(s)
		if err == nil {
			sli = append(sli, x)
		}
	}

	return sli
}

func BubbleSort(sli []int) {
	length := len(sli)

	for i := 0; i < length-1; i++ {
		for j := 0; j < length-i-1; j++ {
			if sli[j] > sli[j+1] {
				Swap(sli, j)
			}
		}
	}
}

func Swap(sli []int, posToSwap int) {
	temp := sli[posToSwap]
	sli[posToSwap] = sli[posToSwap+1]
	sli[posToSwap+1] = temp
}
