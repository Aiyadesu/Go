package main;

import (
	"fmt"
	"math"
)

func Sqrt(value float64) float64 {
	const MAX_ITERATION = 10;
	squareRoot := 1.0;

	for iteration := 0; iteration < MAX_ITERATION; iteration++ {
		squareRoot -= (squareRoot*squareRoot - value) / (2*squareRoot);
	}

	return squareRoot;
}

func mathSqrt(value float64) float64 {
	squareRoot := math.Sqrt(value);

	return squareRoot;
}

func main() {
	for value := 2; value < 100; value++ {
		floatValue := float64(value);
		mySquareRoot := Sqrt(floatValue);
		stlSquareRoot := mathSqrt(floatValue);
		difference := mySquareRoot - stlSquareRoot;

		fmt.Println("For a value of: ", value);
		fmt.Println("My Square Root returns: ", mySquareRoot);
		fmt.Println("Standard Library Square Root returns: ", stlSquareRoot);

		fmt.Println("The difference is: ", difference);
		fmt.Println("");
	}

}
