package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func checkEmptyFile(path string) {
	dirs, err := os.ReadDir(path)

	if err != nil {
		panic(err)
	}

	for _, dir := range dirs {
		nestedPath := filepath.Join(path, dir.Name())

		if dir.IsDir() {
			checkEmptyFile(nestedPath)
			continue
		}

		file, err := os.ReadFile(nestedPath)

		if err != nil {
			panic(err)
		}

		if content := strings.TrimSpace(string(file)); len(content) == 0 {
			fmt.Printf("Empty file: %s\n", nestedPath)
		} else {
			fmt.Printf("File with content: %s\n", nestedPath)
		}
	}

}

func main() {
	checkEmptyFile("./files")
}
