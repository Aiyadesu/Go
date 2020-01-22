package main

import (
	"io"
	"os"
	"strings"
	"fmt"
	"unicode"
	"unicode/utf8"
)

type rot13Reader struct {
	r io.Reader
}

type reader interface {
	Read(io.Reader)
}

func (rotReader rot13Reader) Read(byteArray []byte) (int, error) {
	r := rotReader.r;

	for {
		n, err := r.Read(byteArray);
		b := byteArray[0:n];

		fmt.Printf("n = %v err = %v\n", n, err);
		fmt.Printf("byteArray[:n] = %q\n", b[:n]);

		fmt.Printf("b[:n] = %q\n", b[:n]);

		for len(b) > 0 {
			r, size := utf8.DecodeRune(byteArray);
			fmt.Printf("%c %v\n", r, size);

			b = b[size:];

			if unicode.IsLetter(r) {
				byteValue := utf8.DecodeRune(r);
				// newValue := byteValue;
				// b = newValue;

				fmt.Printf("%c", byteValue);
			}
		}

		// for i, v := range b {
		// 	p := make([]byte, 1);
		// 	p[0] = v;
		// 	ru, _ := utf8.EncodeRune(p);
		// 	fmt.Printf("%v",ru);

		// 	if unicode.IsLetter(ru) {
		// 		newValue := v + 13;
		// 		b[i] = v + 13;
		// 		fmt.Printf("%d, %d, %d, %d\n", i, v, newValue, b[i]);
		// 	}
		// }

		fmt.Printf("b[:n] = %q\n", b[:n]);

		if n == 0 {
			return 0, nil;
		}
	}
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
