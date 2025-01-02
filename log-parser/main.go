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

nextLine:
	for scanner.Scan() {
		line := scanner.Text()

		for _, token := range strings.Fields(line) {

			switch token {
			case INFO:
				logSummary[INFO]++
				continue nextLine // only 1 log level to be matched per line
			case WARN:
				logSummary[WARN]++
				continue nextLine
			case ERROR:
				logSummary[ERROR]++
				continue nextLine
			case FATAL:
				logSummary[FATAL]++
				continue nextLine
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
