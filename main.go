package main

import (
	"time"

	"github.com/Heros-Tempus/pokedexcli/internal/pokeapi"
)

func main() {
	pokeapiClient := pokeapi.NewClient(time.Duration(10*time.Second), time.Duration(5*time.Minute))
	repl(&config{
		client:           pokeapiClient,
		nextLocationsURL: nil,
		prevLocationsURL: nil,
	})
}
