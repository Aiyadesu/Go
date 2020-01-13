package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	var wordMap map[string]int = make(map[string]int);

	wordSlice := strings.Fields(s); // Split the string into words (white-space separated)
	wordCount := len(wordSlice); // Get the length of the slice

	for _, element := range wordSlice { // Iterate over the slice, i.e using the `for`s `range` form
		var elementCount = wordMap[element];

		wordMap[element] = elementCount + 1;
	}

	return wordMap;
}

func main() {
	wc.Test(WordCount)
}
