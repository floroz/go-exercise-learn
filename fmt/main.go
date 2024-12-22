package main

import "fmt"

func readFromCLI() {
	var input string
	fmt.Print("Type something: ")
	fmt.Scanln(&input)
	fmt.Println("You typed: ", input)
}

func main() {
	readFromCLI()
}
