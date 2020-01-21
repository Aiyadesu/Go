package main

import (
	"fmt"
	"io"
	"strings"
	"golang.org/x/tour/reader"
)

type MyReader struct{}

type read interface {
	Read(byte) (int, error)
}

func (MyReader) Read([]byte) (int, error) {
	reader := strings.NewReader("A");

	myByte := make([]byte, 8)

	for {
		n, err := reader.Read(myByte);
		fmt.Printf("n = %v err = %v b = %v\n", n, err, myByte);
		fmt.Printf("B[:n] = %q\n", myByte[:n]);

		if err == io.EOF {
			break;
		}
	}

	return 0, nil;
}
// TODO: Add a Read([]byte) (int, error) method to MyReader.

func main() {
	reader.Validate(MyReader{})
}
