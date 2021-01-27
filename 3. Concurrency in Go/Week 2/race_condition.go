package main

import "fmt"

/*
Write two goroutines which have a race condition when executed concurrently.
Explain what the race condition is and how it can occur.
*/

func main() {
	a := 0

	for i := 0; i < 100; i++ {
		go func() { a++ }()
	}

	fmt.Printf("%d", a)
}

/*
Race Condition: condition that occurs when there are conflicting accesses to a resource,
this can lead to different results in multiple executions.

Occurs when multiple process can access and change the data at the same time.
*/
