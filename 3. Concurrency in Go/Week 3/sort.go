package main

import (
	"fmt"
	"sort"
	"strconv"
)

/*
Write a program to sort an array of integers.
The program should partition the array into 4 parts, each of which is sorted by a different goroutine.
Each partition should be of approximately equal size.
Then the main goroutine should merge the 4 sorted subarrays into one large sorted array.

The program should prompt the user to input a series of integers.
Each goroutine which sorts ¼ of the array should print the subarray that it will sort.
When sorting is complete, the main goroutine should print the entire sorted list.
*/

func ReadSliceFromInput() []int {
	var s string
	var sli = make([]int, 0, 12)

	fmt.Println("Enter 12 integers")
	for i := 0; i < 12; i++ {
		_, _ = fmt.Scan(&s)
		x, err := strconv.Atoi(s)
		if err == nil {
			sli = append(sli, x)
		}
	}

	return sli
}

func sortSlice(ints []int, c chan []int) {
	sort.Ints(ints)
	fmt.Println(ints)
	c <- ints
}

func merge(arr1, arr2 []int) []int {
	finalSlice := make([]int, len(arr1)+len(arr2))

	i := 0
	for len(arr1) > 0 && len(arr2) > 0 {
		if arr1[0] < arr2[0] {
			finalSlice[i] = arr1[0]
			arr1 = arr1[1:]
		} else {
			finalSlice[i] = arr2[0]
			arr2 = arr2[1:]
		}
		i++
	}

	for j := 0; j < len(arr1); j++ {
		finalSlice[i] = arr1[j]
		i++
	}
	for j := 0; j < len(arr2); j++ {
		finalSlice[i] = arr2[j]
		i++
	}
	return finalSlice
}

func main() {
	sli := ReadSliceFromInput()

	block1End := len(sli) / 4
	block2End := 2 * len(sli) / 4
	block3End := 3 * len(sli) / 4

	c := make(chan []int, 4)
	go sortSlice(sli[:block1End], c)
	go sortSlice(sli[block1End:block2End], c)
	go sortSlice(sli[block2End:block3End], c)
	go sortSlice(sli[block3End:], c)

	arr1 := <-c
	arr2 := <-c
	arr3 := <-c
	arr4 := <-c

	finalSlice := merge(merge(arr1, arr2), merge(arr3, arr4))

	fmt.Println(finalSlice)
}
