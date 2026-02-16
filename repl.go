package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Heros-Tempus/pokedexcli/internal/pokeapi"
)

func repl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			break
		}
		input := cleanInput(scanner.Text())
		if len(input) == 0 {
			continue
		}
		command, exists := getCommands()[input[0]]
		if !exists {
			fmt.Println("Unknown command")
			continue
		}
		err := command.callback(cfg, input[1:]...)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func cleanInput(text string) []string {
	lowercase := strings.ToLower(text)
	return strings.Fields(lowercase)
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	var commands map[string]cliCommand = map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Display help information",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Display map information",
			callback:    getLocations,
		},
		"mapb": {
			name:        "mapb",
			description: "Display previous map information",
			callback:    getPreviousLocations,
		},
		"explore": {
			name:        "explore",
			description: "Explore a location",
			callback:    exploreLocation,
		},
		"catch": {
			name:        "catch",
			description: "Attempt to catch a Pokemon",
			callback:    catchPokemon,
		},
	}
	return commands
}

type config struct {
	client           pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
	caughtPokemon    map[string]pokeapi.Pokemon
}
