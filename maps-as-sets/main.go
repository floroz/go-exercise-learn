package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	wordSet := map[string]bool{}

	file, err := os.Open("text.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		word := scanner.Text()

		if !wordSet[word] {
			wordSet[word] = true
		}
	}

	fmt.Printf("Found %d unique words.\n", len(wordSet))

}
