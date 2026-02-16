package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) CatchRate(species string) (CatchRate, error) {
	url := baseURL + "/pokemon-species/" + species

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return CatchRate{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return CatchRate{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 299 && resp.StatusCode < 500 {
		return CatchRate{}, fmt.Errorf("Client error: %v, make sure \"%s\" is a valid species name", resp.StatusCode, species)
	} else if resp.StatusCode > 499 {
		return CatchRate{}, fmt.Errorf("Server error: %v", resp.StatusCode)
	}

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return CatchRate{}, err
	}

	rate := CatchRate{}
	err = json.Unmarshal(dat, &rate)
	if err != nil {
		return CatchRate{}, err
	}

	return rate, nil
}

func (c *Client) GetPokemon(species string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + species

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 299 && resp.StatusCode < 500 {
		return Pokemon{}, fmt.Errorf("Client error: %v, make sure \"%s\" is a valid species name", resp.StatusCode, species)
	} else if resp.StatusCode > 499 {
		return Pokemon{}, fmt.Errorf("Server error: %v", resp.StatusCode)
	}

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}

	pokemon := Pokemon{}
	err = json.Unmarshal(dat, &pokemon)
	if err != nil {
		return Pokemon{}, err
	}

	return pokemon, nil
}
