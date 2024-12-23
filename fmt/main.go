package main

import "fmt"

// fmt.Scan functions allow to read from stdin
func readFromStdin() {
	var input string
	fmt.Print("Type something: ")
	fmt.Scanln(&input)
	fmt.Printf("You typed: %s", input)
}

// %T prints the type of the variable
func printType() {
	str := "foo"
	digit := 10
	float := 2.2

	fmt.Printf("str: %T - digit: %T - float: %T\n", str, digit, float)
}

// %[<index>]<symbol> syntax allow to reference previous values in the template - index starts at 1;
func referToPreviousTemplateValue() {
	a, b := "Pizza", "Burger"

	fmt.Printf("I ate a %v, then a %v, but was hungry so ate another %[1]v", a, b)
}

// You can control precision output of float verb in formatter
func floatPrintingPrecision() {
	balance := 567.3432

	fmt.Printf("Your balance is - %f\n", balance)
	fmt.Printf("Your balance is - %.0f\n", balance)
	fmt.Printf("Your balance is - %.1f\n", balance)
	fmt.Printf("Your balance is - %.3f\n", balance)
}

func main() {
	// readFromStdin()
	// printType()
	// referToPreviousTemplateValue()
	// floatPrintingPrecision()
}
