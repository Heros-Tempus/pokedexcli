package main

import (
	"fmt"
)

func listCaughtPokemon(cfg *config, _ ...string) error {
	if len(cfg.caughtPokemon) == 0 {
		fmt.Println("You haven't caught any Pokemon yet!")
		return nil
	}
	fmt.Println("Your Pokedex:")
	for name := range cfg.caughtPokemon {
		fmt.Printf("  - %s\n", name)
	}
	return nil
}
