package main

import (
	"strings"
	"fmt"
	"bufio"
	"os"
)

func startRepl() {
	input := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		input.Scan()

		inputString := input.Text()
		cleanString := cleanInput(inputString)

		if len(cleanString) == 0 {
			continue
		}

		firstWord := cleanString[0]
		fmt.Println("Your command was:", firstWord)
	}
}

func cleanInput(text string) []string {
	text = strings.ToLower(text)
	splittedSlice := strings.Fields(text)

	return splittedSlice
}