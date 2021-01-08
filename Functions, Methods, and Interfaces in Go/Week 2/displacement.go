package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	sli := ReadSliceFromInput()

	acceleration := sli[0]
	initialVelocity := sli[1]
	initialDisplacement := sli[2]

	fn := GenDisplaceFn(acceleration, initialVelocity, initialDisplacement)

	fmt.Println("Enter time: ")
	var s string
	_, _ = fmt.Scan(&s)
	x, err := strconv.ParseFloat(s, 64)
	if err == nil {
		finalPosition := fn(x)
		fmt.Printf("Your final position is: %f", finalPosition)
	}
}

func GenDisplaceFn(acceleration, initialVelocity, initialDisplacement float64) func(float64) float64 {
	return func(time float64) float64 {
		return acceleration/2*math.Pow(time, 2) +
			initialVelocity*time +
			initialDisplacement
	}
}

func ReadSliceFromInput() []float64 {
	var s string
	var sli = make([]float64, 0, 3)

	messages := [3]string{"Acceleration", "Initial Velocity", "Initial Displacement"}

	for i := 0; i < 3; i++ {
		fmt.Printf("Enter your %s:", messages[i])
		fmt.Println()
		_, _ = fmt.Scan(&s)
		x, err := strconv.ParseFloat(s, 64)

		if err == nil {
			sli = append(sli, x)
		}
	}

	return sli
}
