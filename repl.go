package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	input := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		input.Scan()

		inputString := input.Text()
		commandName := cleanInput(inputString)

		if len(commandName) == 0 {
			continue
		}

		if command, ok := getCommands()[commandName[0]]; !ok {
			fmt.Println("Unknown command")
			continue
		} else {
			config := config{next: "", previous: ""}
			err := command.callback(&config)
			if err != nil {
				fmt.Println("Error executing command:", command, err)
			}
		}
	}
}

func cleanInput(text string) []string {
	text = strings.ToLower(text)
	splittedSlice := strings.Fields(text)

	return splittedSlice
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	next	 string
	previous string
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the next 20 locations of Pokemon World",
			callback:    commandHelp,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 locations of Pokemon World",
			callback:    commandHelp,
		},
	}
}
