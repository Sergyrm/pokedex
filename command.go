package main

import (
	"errors"
	"fmt"
	"os"
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
	location, err := conf.pokeapiClient.GetLocationAreas(conf.next)
	if err != nil {
		return err
	}

	conf.next = location.Next
	conf.previous = location.Previous

	for _, result := range location.Results {
		fmt.Printf("%s\n", result.Name)
	}

	return nil
}

func commandMapb(conf *config) error {
	if conf.previous == nil {
		return errors.New("you're on the first page")
	}
	location, err := conf.pokeapiClient.GetLocationAreas(conf.previous)
	if err != nil {
		return err
	}

	conf.next = location.Next
	conf.previous = location.Previous

	for _, result := range location.Results {
		fmt.Printf("%s\n", result.Name)
	}

	return nil
}
