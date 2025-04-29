package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	input := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		input.Scan()

		inputString := input.Text()
		cleanString := cleanInput(inputString)
		firstWord := cleanString[0]
		fmt.Println("Your command was:", firstWord)
	}
}
