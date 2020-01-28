package main;

import (
	"fmt"
)

type ErrNegativeSqrt float64;

type error interface {
	Error() string
}

func (err ErrNegativeSqrt) Error() string {
	return fmt.Sprint("cannot Sqrt negative number: %v", float64(err));
}

func Sqrt(value float64) (float64, error) {
	if value < 0 {
		return value, ErrNegativeSqrt(value);
	}

	const MAX_ITERATION = 10;
	squareRoot := 1.0;

	for iteration := 0; iteration < MAX_ITERATION; iteration++ {
		squareRoot -= (squareRoot*squareRoot - value) / (2*squareRoot);
	}

	return squareRoot, nil;
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
