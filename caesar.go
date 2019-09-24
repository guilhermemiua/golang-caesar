package main

import (
	"fmt"
	"strings"
)

func main() {
	encoder("!hello world", 3)
}

func encoder(word string, leap int) {
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	word = strings.ToLower(word)
	var cipher string = ""

	for i := 0; i < len(word); i++ {
		var letter string = string(word[i])
		var index int = strings.Index(alphabet, letter)
		var position int = index + leap

		// Verify if it is a letter not included in alphabet
		if index < 0 {
			cipher = cipher + string(word[i])
			continue
		}

		// Verify if index + leap passed the alphabet index
		if position > 25 {
			position = position - 25
		}

		cipher = cipher + string(alphabet[position])
	}

	fmt.Println(word)
	fmt.Println(cipher)
}
