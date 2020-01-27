package main

import (
	"golang.org/x/tour/reader"
)

type MyReader struct{}

type read interface {
	Read(byte) (int, error)
}

// Refactored to use 'AntoineAugsti's solution: https://github.com/AntoineAugusti/go-exercices/blob/master/7-exercise-reader.go
func (MyReader) Read(byteArray []byte) (int, error) {
	byteArray = byteArray[:cap(byteArray)];

	for index := range byteArray {
		byteArray[index] = 'A';
	}

	return cap(byteArray), nil;
}

func main() {
	reader.Validate(MyReader{})
}
