package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

func main() {
	for {
		var guess string

		fmt.Print("Guess your lucky number between 0 and 100: ")
		fmt.Scanln(&guess)

		num, err := strconv.Atoi(guess)

		if err != nil {
			fmt.Println("Invalid guess - please type a number.")
			continue
		}

		randomInt := rand.Intn(100)

		if randomInt == num {
			fmt.Println("Correct!")
		} else {
			fmt.Printf("Close enough...it was %d\n", randomInt)
		}
	}
}
