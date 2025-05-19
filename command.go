package main

import (
	"errors"
	"fmt"
	"os"
	"math/rand"
)

func commandExit(conf *config, params []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(conf *config, params []string) error {
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

func commandMap(conf *config, params []string) error {
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

func commandMapb(conf *config, params []string) error {
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

func commandExplore(conf *config, params []string) error {
	if len(params) < 1 {
		return errors.New("please provide a location")
	}

	fmt.Printf("Exploring %s...\n", params[0])
	locationDetails, err := conf.pokeapiClient.GetPokemonByLocation(conf.previous, params[0])
	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")
	for _, result := range locationDetails.PokemonEncounters {
		fmt.Printf(" - %s\n", result.Pokemon.Name)
	}

	return nil
}

func commandCatch(conf *config, params []string) error {
	if len(params) < 1 {
		return errors.New("please provide a pokemon")
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", params[0])
	pokemon, err := conf.pokeapiClient.GetPokemonInfo(conf.previous, params[0])
	if err != nil {
		return err
	}

	escapeRate := rand.Intn(pokemon.BaseExperience)

	if escapeRate <= 40 {
		fmt.Printf("%s was caught!\n", pokemon.Name)
	} else {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	}

	return nil
}