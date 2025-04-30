package main

import (
	"fmt"
	"os"
	"github.com/Sergyrm/pokedex/internal/pokeapi"
)

func commandExit(conf *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(conf *config) error {
	_, err := fmt.Println("Welcome to the Pokedex!")
	if err != nil {
		return err
	}

	_, err = fmt.Println("Usage:")
	if err != nil {
		return err
	}

	for _, command := range getCommands() {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}

	return nil
}

func commandMap(conf *config) error {
	location := getPokeApiData(conf.next)
	fmt.Println(location)
	return nil
}
