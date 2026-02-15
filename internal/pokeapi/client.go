package pokeapi

import (
	"net/http"
	"time"

	"github.com/Heros-Tempus/pokedexcli/internal/pokecache"
)

type Client struct {
	httpClient http.Client
	cache      pokecache.Cache
}

func NewClient(timeout, cacheDuration time.Duration) Client {
	return Client{httpClient: http.Client{Timeout: timeout}, cache: pokecache.NewCache(cacheDuration)}
}

type LocationListResp struct {
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}
