package main

import (
	"fmt"
	"strings"
)

func main() {
	encipher("!hello world", 3)
}

func encipher(word string, leap int) {
	letters := "abcdefghijklmnopqrstuvwxyz"
	word = strings.ToLower(word)
	var cipher string = ""

	for i := 0; i < len(word); i++ {
		var letter string = string(word[i])
		var index int = strings.Index(letters, letter)
		var position int = index + leap

		if index < 0 {
			cipher = cipher + string(word[i])
			continue
		}

		if position > 25 {
			position = position - 25
		}

		cipher = cipher + string(letters[position])
	}

	fmt.Println(word)
	fmt.Println(cipher)
}
