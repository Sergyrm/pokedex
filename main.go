package main

import (
	"time"

	"github.com/Sergyrm/pokedex/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	conf := &config{
		pokeapiClient: pokeClient,
	}

	startRepl(conf)
}
