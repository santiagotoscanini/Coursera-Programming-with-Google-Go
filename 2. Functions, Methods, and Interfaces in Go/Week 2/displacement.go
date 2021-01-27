package main

import (
	"fmt"
	"math"
	"strconv"
)

/*
Let us assume the following formula for displacement s as a function of time t, acceleration a, initial velocity vo, and initial displacement so.
s =½ a t2 + vot + so

Write a program which first prompts the user to enter values for acceleration, initial velocity, and initial displacement.
Then the program should prompt the user to enter a value for time and the program should compute the displacement after the entered time.

You will need to define and use a function called GenDisplaceFn() which takes three float64 arguments, acceleration a, initial velocity vo, and initial displacement so.
GenDisplaceFn() should return a function which computes displacement as a function of time, assuming the given values acceleration, initial velocity, and initial displacement.
The function returned by GenDisplaceFn() should take one float64 argument t, representing time, and return one float64 argument which is the displacement travelled after time t.

For example, let’s say that I want to assume the following values for acceleration, initial velocity, and initial displacement: a = 10, vo = 2, so = 1.
I can use the following statement to call GenDisplaceFn() to generate a function fn which will compute displacement as a function of time.

fn := GenDisplaceFn(10, 2, 1)

Then I can use the following statement to print the displacement after 3 seconds.

fmt.Println(fn(3))

And I can use the following statement to print the displacement after 5 seconds.

fmt.Println(fn(5))`
*/

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
