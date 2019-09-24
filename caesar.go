package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Phrase used: Hello World!")
	encoder("Hello World!", 3)
	decoder("Khoor Zruog!", 3)
}

func encoder(phrase string, leap int) {
	alphabetLower := "abcdefghijklmnopqrstuvwxyz"
	alphabetUpper := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var cipher string = ""

	for i := 0; i < len(phrase); i++ {
		var letter string = string(phrase[i])
		var index int
		var lower bool = strings.Contains(alphabetLower, letter)

		if lower {
			index = strings.Index(alphabetLower, letter)
		} else {
			index = strings.Index(alphabetUpper, letter)
		}

		var position int = index + leap

		// Verify if it is a letter not included in alphabet
		if index < 0 {
			cipher = cipher + string(phrase[i])
			continue
		}

		// Verify if index + leap passed the alphabet index
		if position > 25 {
			position = position - 25
		}

		if lower {
			cipher = cipher + string(alphabetLower[position])
		} else {
			cipher = cipher + string(alphabetUpper[position])
		}
	}

	fmt.Println("Cipher: " + cipher)
}

func decoder(phrase string, leap int) {
	alphabetLower := "abcdefghijklmnopqrstuvwxyz"
	alphabetUpper := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var answer string = ""

	for i := 0; i < len(phrase); i++ {
		var letter string = string(phrase[i])
		var lower bool = strings.Contains(alphabetLower, letter)
		var index int

		if lower {
			index = strings.Index(alphabetLower, letter)
		} else {
			index = strings.Index(alphabetUpper, letter)
		}
		var position int = index - leap

		// Verify if it is a letter not included in alphabet
		if index < 0 {
			answer = answer + string(phrase[i])
			continue
		}

		if position < 0 {
			position = 26 + position
		}

		if lower {
			answer = answer + string(alphabetLower[position])
		} else {
			answer = answer + string(alphabetUpper[position])
		}
	}

	fmt.Println("Decoded phrase: " + answer)
}
