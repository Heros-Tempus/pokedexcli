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
