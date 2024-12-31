package main

import (
	"fmt"
	"os"
)

var spamDomainName = []byte("spamlink.com")

const prefixLength = len("http://")

func hasHttpPrefix(file *[]byte, i int) bool {
	return i+prefixLength <= len(*file) && string((*file)[i:i+prefixLength]) == "http://"
}

func hasSpamDomain(file *[]byte, j int) bool {
	if j+len(spamDomainName) > len(*file) {
		return false
	}
	for i := range spamDomainName {
		if (*file)[j+i] != spamDomainName[i] {
			return false
		}
	}
	return true
}

func isEndOfURI(file *[]byte, j int) bool {
	if j >= len(*file) {
		return true
	}
	curr := (*file)[j]
	next := byte(' ')
	if j+1 < len(*file) {
		next = (*file)[j+1]
	}
	return curr == ' ' || curr == '\n' || (curr == '.' && next == ' ')
}

func SpamMasker(file *[]byte) {
	text := *file
	size := len(text)

	fmt.Printf("%s\n", text)
	fmt.Println("________________________________________________________")

	// the buffer where to save the new string as we process each byte
	buff := make([]byte, 0, size)

	for i := 0; i < size; i++ {
		if hasHttpPrefix(&text, i) {
			// we move the j position after the `http://` to inspect the domain
			j := i + prefixLength

			// out of bounds
			if j >= size {
				buff = append(buff, text[i])
				continue
			}

			if !hasSpamDomain(&text, j) {
				buff = append(buff, text[i])
				continue
			}

			// we have already processed this portion so we can append it to the buffer
			buff = append(buff, []byte("http://")...)

			for !isEndOfURI(&text, j) {
				buff = append(buff, byte('*'))
				j++
			}

			// fast forward the loop since all bytes until end have been processed already
			i = j - 1

			continue
		}

		buff = append(buff, text[i])
	}

	fmt.Printf("%s", buff)
}

func main() {
	file, _ := os.ReadFile("text.txt")

	SpamMasker(&file)
}
