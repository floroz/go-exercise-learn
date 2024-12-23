package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// string should be counted by bytes, using len, or by runes
func countString() {
	// count bytes
	fmt.Println(len(`a 🙃`))

	// count runes
	// these two are equivalents
	fmt.Println(utf8.RuneCountInString(`a 🙃`))
	fmt.Println([]rune(`a 🙃`))
}

// string package has built in methods to manipulate like repeat
func getInputAndRepeatString() {
	bangs := strings.Repeat("!", 3)

	var t string
	fmt.Scanln(&t)

	fmt.Println(strings.ToUpper(t) + bangs)

}

// iota autogenerates number fo constant
func numberGeneratorIota() {
	const (
		monday = iota + 1
		tuesday
		wednesday
		thursday
		friday
		saturday
		sunday
	)

	fmt.Println(monday, tuesday, wednesday, thursday, friday, saturday, sunday)
}

func main() {
	// countString()
	// getInputAndRepeatString()
	// numberGeneratorIota()
}
