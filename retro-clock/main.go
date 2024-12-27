package main

import (
	"fmt"
	"time"

	"github.com/inancgumus/screen"
)

type clockSymbol [5][3]bool

var clockSymbolsMap = map[string]clockSymbol{
	"0": {
		{true, true, true},
		{true, false, true},
		{true, false, true},
		{true, false, true},
		{true, true, true},
	},
	"1": {
		{false, true, false},
		{true, true, false},
		{false, true, false},
		{false, true, false},
		{true, true, true},
	},
	"2": {
		{true, true, true},
		{false, false, true},
		{true, true, true},
		{true, false, false},
		{true, true, true},
	},
	"3": {
		{true, true, true},
		{false, false, true},
		{true, true, true},
		{false, false, true},
		{true, true, true},
	},
	"4": {
		{true, false, true},
		{true, false, true},
		{true, true, true},
		{false, false, true},
		{false, false, true},
	},
	"5": {
		{true, true, true},
		{true, false, false},
		{true, true, true},
		{false, false, true},
		{true, true, true},
	},
	"6": {
		{true, true, true},
		{true, false, false},
		{true, true, true},
		{true, false, true},
		{true, true, true},
	},
	"7": {
		{true, true, true},
		{false, false, true},
		{false, false, true},
		{false, true, false},
		{false, true, false},
	},
	"8": {
		{true, true, true},
		{true, false, true},
		{true, true, true},
		{true, false, true},
		{true, true, true},
	},
	"9": {
		{true, true, true},
		{true, false, true},
		{true, true, true},
		{false, false, true},
		{true, true, true},
	},
	":": {
		{false, false, false},
		{false, true, false},
		{false, false, false},
		{false, true, false},
		{false, false, false},
	},
}

func main() {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	// Run every second
	for now := range ticker.C {
		screen.Clear()
		screen.MoveTopLeft()

		h := fmt.Sprintf("%02d", now.Hour())
		m := fmt.Sprintf("%02d", now.Minute())
		s := fmt.Sprintf("%02d", now.Second())

		timeSequence := h + ":" + m + ":" + s

		var time []clockSymbol

		// Convert each character in the time sequence to its corresponding clock symbol
		for _, digit := range timeSequence {
			time = append(time, clockSymbolsMap[string(digit)])
		}

		// Print the time row by row across digits
		for row := 0; row < 5; row++ {
			fmt.Print("\n")
			// Loop through all columns of all digits and print them
			// Use a space to separate the digits
			// Print the value with "true" as a green filled box
			for _, digit := range time {
				for col := 0; col < 3; col++ {
					if digit[row][col] {
						fmt.Print("\033[92m██\033[0m") // Fluorescent green
					} else {
						fmt.Print("  ")
					}
				}
				fmt.Print("  ") // Separate the digits
			}
		}
		fmt.Print("\n")
	}
}
