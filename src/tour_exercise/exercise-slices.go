package main

import "fmt";
import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	fmt.Println(dx);
	fmt.Println(dy);

	s := []uint8{2, 3, 5, 7, 11, 13};
  a := []uint8{2, 5, 8, 9}
	testValue := [][]uint8{s, a};
	return testValue;
}

func main() {
	pic.Show(Pic)
}
