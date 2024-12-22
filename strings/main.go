package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	fmt.Println(len(`a 🙃`))
	fmt.Println(utf8.RuneCountInString(`a 🙃`))
}
