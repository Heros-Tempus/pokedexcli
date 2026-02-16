package main

import (
	"fmt"
	"math"
	"math/rand/v2"
	"strings"
	"unicode"
)

func catchPokemon(cfg *config, args ...string) error {
	if len(args) < 1 {
		return fmt.Errorf("Usage: catch <pokemon_name> (optional)<ball_type>")
	}
	pokemonName := args[0]
	ballMultiplier := 1.0
	if len(args) > 1 {
		ballType := args[1]
		switch ballType {
		case "great", "greatball":
			ballMultiplier = 1.5
		case "ultra", "ultraball":
			ballMultiplier = 2.0
		case "master", "masterball":
			ballMultiplier = 255.0
		}
	}
	CatchRate, err := cfg.client.CatchRate(pokemonName)
	if err != nil {
		return err
	}
	pokemon, err := cfg.client.GetPokemon(pokemonName)
	if err != nil {
		return err
	}
	captureRate := CatchRate.CaptureRate

	if ballMultiplier != 1.0 {
		fmt.Printf("Throwing a %sball at %s...\n", CapitalizeFirst(args[1]), pokemonName)
	} else {
		fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)
	}

	if Catch(captureRate, ballMultiplier) {
		fmt.Printf("%s was caught!\n", pokemonName)
		cfg.caughtPokemon[pokemonName] = pokemon
	} else {
		fmt.Printf("%s escaped!\n", pokemonName)
	}

	return nil
}

func CapitalizeFirst(s string) string {
	if len(s) == 0 {
		return s
	}

	runes := []rune(strings.ToLower(s))
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}
func Catch(catchRate int, ballMultiplier float64) bool {
	x := (float64(catchRate) * ballMultiplier) / 3.0
	if x >= 255 {
		return true
	}
	y := math.Floor(65536.0 / math.Pow(255.0/x, 3.0/16.0))
	probability := math.Pow(y/65536.0, 4)
	return probability > rand.Float64()
}
