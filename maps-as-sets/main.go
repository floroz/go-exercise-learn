package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	wordSet := map[string]uint8{}

	file, err := os.Open("text.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		for _, token := range strings.Fields(line) {
			if _, ok := wordSet[token]; !ok {
				wordSet[token]++
			}
		}
	}
	fmt.Println(wordSet)

	fmt.Printf("%d\n", len(wordSet))

}
