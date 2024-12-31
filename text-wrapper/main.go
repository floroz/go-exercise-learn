package main

import (
	"fmt"
	"os"
	"unicode"
)

const (
	maxWidth = 40
)

func main() {
	file, _ := os.ReadFile("text.txt")
	fmt.Println("___________________________________BEFORE__________________________")
	fmt.Printf("%s\n", file)
	fmt.Println("__________________________________________________________________________")

	buffer := make([]rune, 0, len(file))

	width := 0

	for _, r := range string(file) {
		switch {
		// when we hit a new line we reset the counter
		case r == '\n':
			buffer = append(buffer, r)
			width = 1
		// if we hit the max width - we have to wrap after the end of the word
		// it wont be a new line since the case is handled before
		case width >= maxWidth && unicode.IsSpace(r):
			buffer = append(buffer, '\n')
			width = 1
		default:
			buffer = append(buffer, r)
			width++
		}

	}

	fmt.Printf("%s\n", string(buffer))
	fmt.Println("__________________________________AFTER____________________________")

}
