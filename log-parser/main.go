package main

import (
	"bufio"
	"fmt"
	"os"
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
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		token := scanner.Text()

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

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println("Logs Summary: ")
	for k, v := range logSummary {
		fmt.Printf("%s: %d\n", k, v)
	}
}
