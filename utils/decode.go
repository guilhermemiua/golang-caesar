package utils

import (
	"fmt"
	"strings"
)

func Decode(phrase string, leap int) {
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
