package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	fibonacciSequence := 0;
	numberToAdd := 0;

	return func() int {
		if numberToAdd == 0 {
			numberToAdd = 1;

			return fibonacciSequence;
		}

		tempValue := fibonacciSequence;
		fibonacciSequence = fibonacciSequence + numberToAdd;
		numberToAdd = tempValue;

		return fibonacciSequence;
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
