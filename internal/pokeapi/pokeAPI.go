package pokeapi

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/Sergyrm/pokedex/internal/pokecache"
)

const (
	baseURL = "https://pokeapi.co/api/v2"
)

type Location struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type Client struct {
	httpClient http.Client
	pokeCache  *pokecache.Cache
}

func NewClient(timeout time.Duration) Client {
	cache := pokecache.NewCache(timeout)

	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		pokeCache: &cache,
	}
}

func (c *Client) GetLocationAreas(pageURL *string) (Location, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	if cachedData, found := c.pokeCache.Get(url); found {
		location := Location{}
		err := json.Unmarshal(cachedData, &location)
		if err == nil {
			return location, nil
		}
	}

	res, err := http.Get(url)
	if err != nil {
		return Location{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		return Location{}, err
	}

	c.pokeCache.Add(url, body)

	location := Location{}
	err = json.Unmarshal(body, &location)
	if err != nil {
		return Location{}, err
	}

	return location, nil
}
