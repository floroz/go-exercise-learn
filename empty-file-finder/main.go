package main

import (
	"fmt"
	"os"
	"path/filepath"
)

const outputFilename = "output.txt"

func checkEmptyFile(path string, outputFile *os.File) {
	// read dirs
	dirs, err := os.ReadDir(path)

	if err != nil {
		panic(err)
	}

	// no files or directory - early exit
	if len(dirs) == 0 {
		return
	}

	for _, dir := range dirs {
		nestedPath := filepath.Join(path, dir.Name())

		// a directory
		if dir.IsDir() {
			checkEmptyFile(nestedPath, outputFile)
			continue
		}
		// a file
		fileInfo, err := dir.Info()

		if err != nil {
			panic(err)
		}

		if fileInfo.Size() == 0 {
			outputFile.WriteString(fmt.Sprintf("Empty file: %s\n", nestedPath))
		}

	}

}

func main() {
	// prepare the output file
	err := os.WriteFile(outputFilename, []byte(nil), 0644)

	if err != nil {
		panic("failed to clear the output file.")
	}

	// keep file open to speed up writes
	outputFile, err := os.OpenFile(outputFilename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		panic("failed to open the output file.")
	}

	defer outputFile.Close()

	checkEmptyFile("./files", outputFile)
}
