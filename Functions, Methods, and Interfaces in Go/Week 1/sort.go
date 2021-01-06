package main

import (
	"fmt"
	"strconv"
)

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
