package main

import (
	"fmt"
	"golang-caesar/utils"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1]
	phrase := os.Args[2]
	leap, err := strconv.Atoi(os.Args[3])

	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	if args == "encode" {
		utils.Encode(phrase, leap)
	} else if args == "decode" {
		utils.Decode(phrase, leap)
	} else {
		fmt.Println("Function not available.")
	}

	os.Exit(0)
}
