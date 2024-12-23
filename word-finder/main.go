package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	text := readInput("Insert text to search from:\nFor example: \"Earth is the only planet with living beings\"")
	if text == "" {
		fmt.Println("Invalid text")
		return
	}

	fields := strings.Fields(strings.ToLower(text))
	if len(fields) == 0 {
		fmt.Println("Invalid text")
		return
	}

	query := readInput("Add search terms separated by comma\nFor example: earth,planet")
	if query == "" {
		fmt.Println("Invalid string")
		return
	}

	searchTerms := strings.Split(strings.ToLower(strings.TrimSpace(query)), ",")
	matches := findMatches(fields, searchTerms)

	if len(matches) == 0 {
		fmt.Println("No matches found.")
		return
	}

	for _, m := range matches {
		fmt.Printf("Match found for `%s` search term\n", m)
	}
}

func readInput(prompt string) string {
	fmt.Println(prompt)
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		return ""
	}
	return strings.TrimSpace(input)
}

func findMatches(fields, searchTerms []string) []string {
	var matches []string
	for _, st := range searchTerms {
		if slices.Contains(fields, st) {
			matches = append(matches, st)
		}
	}
	return matches
}
