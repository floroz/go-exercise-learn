package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	INFO  = "INFO"
	WARN  = "WARN"
	ERROR = "ERROR"
	FATAL = "FATAL"
)

func main() {

	logSummary := map[string]int{
		INFO:  0,
		WARN:  0,
		ERROR: 0,
		FATAL: 0,
	}

	file, err := os.Open("logs.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		for _, token := range strings.Fields(line) {

			switch token {
			case INFO:
				logSummary[INFO]++
			case WARN:
				logSummary[WARN]++
			case ERROR:
				logSummary[ERROR]++
			case FATAL:
				logSummary[FATAL]++
			}
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println("Logs Summary: ")
	for k, v := range logSummary {
		fmt.Printf("%s: %d\n", k, v)
	}
}
