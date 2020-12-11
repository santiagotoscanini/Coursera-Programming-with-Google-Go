package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	var input string
	_, _ = fmt.Scan(&input)

	slice := make([]int, 0)
	for input != "X" {
		intFormat, _ := strconv.Atoi(input)

		slice = append(slice, intFormat)
		sort.Slice(slice, func(i, j int) bool {
			return slice[i] < slice[j]
		})

		fmt.Println(slice)
		_, _ = fmt.Scan(&input)
	}
}
