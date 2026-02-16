package main

import (
	"fmt"
)

func exploreLocation(cfg *config, args ...string) error {
	if len(args) < 1 {
		return fmt.Errorf("Usage: explore <location-name>")
	}
	locationName := args[0]
	resp, err := cfg.client.ExploreLocation(locationName)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\nFound Pokemon:\n", locationName)
	for _, encounter := range resp.PokemonEncounters {
		fmt.Println("- " + encounter.Pokemon.Name)
	}
	return nil
}
