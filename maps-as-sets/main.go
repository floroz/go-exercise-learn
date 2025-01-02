package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	wordSet := map[string]uint8{}

	file, err := os.Open("text.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		token := scanner.Text()

		if _, ok := wordSet[token]; !ok {
			wordSet[token]++
		}
	}

	fmt.Printf("Found %d unique words.\n", len(wordSet))

}
