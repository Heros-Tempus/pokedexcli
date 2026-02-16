package main

import (
	"fmt"
	"math"
	"math/rand/v2"

	"github.com/Heros-Tempus/pokedexcli/internal/pokeapi"
)

func catchPokemon(cfg *config, args ...string) error {
	if len(args) < 1 {
		return fmt.Errorf("Usage: catch <pokemon_name> [ball_type]")
	}

	pokemonName := args[0]
	if _, ok := cfg.caughtPokemon[pokemonName]; ok {
		fmt.Printf("You have already caught %s!\n", pokemonName)
		return nil
	}

	type rateRes struct {
		rate int
		err  error
	}
	type pokeRes struct {
		pokemon pokeapi.Pokemon
		err     error
	}
	rChan := make(chan rateRes, 1)
	pChan := make(chan pokeRes, 1)
	go func() {
		rate, err := cfg.client.CatchRate(pokemonName)
		if err != nil {
			rChan <- rateRes{0, err}
			return
		}
		rChan <- rateRes{rate: rate.CaptureRate, err: nil}
	}()
	go func() {
		pokemon, err := cfg.client.GetPokemon(pokemonName)
		if err != nil {
			pChan <- pokeRes{err: err}
			return
		}
		pChan <- pokeRes{pokemon: pokemon, err: nil}
	}()

	ballMultiplier := 1.0
	displayBall := "Poke"
	if len(args) > 1 {
		ballType := args[1]
		switch ballType {
		case "great", "greatball":
			ballMultiplier = 1.5
			displayBall = "Great"
		case "ultra", "ultraball":
			ballMultiplier = 2.0
			displayBall = "Ultra"
		case "master", "masterball":
			ballMultiplier = 255.0
			displayBall = "Master"
		}
	}
	rRes := <-rChan
	if rRes.err != nil {
		return fmt.Errorf("Failed to get capture rate: %w.", rRes.err)
	}
	pRes := <-pChan
	if pRes.err != nil {
		return fmt.Errorf("Failed to get pokemon: %w.", pRes.err)
	}

	captureRate := rRes.rate
	pokemon := pRes.pokemon

	fmt.Printf("Throwing a %sball at %s...\n", displayBall, pokemonName)

	if Catch(captureRate, ballMultiplier) {
		fmt.Printf("%s was caught!\n", pokemonName)
		cfg.caughtPokemon[pokemonName] = pokemon
	} else {
		fmt.Printf("%s escaped!\n", pokemonName)
	}

	return nil
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
