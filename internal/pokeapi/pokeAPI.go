package pokeapi

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)

const (
	baseURL = "https://pokeapi.co/api/v2"
)

type Location struct {
	Count    	int    `json:"count"`
	Next    	*string `json:"next"`
	Previous	*string `json:"previous"`
	Results  	[]struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type Client struct {
	httpClient http.Client
}

func NewClient(timeout time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}

func (c *Client) GetLocationAreas(pageURL *string) (Location, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
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

	location := Location{}
	err = json.Unmarshal(body, &location)
	if err != nil {
		return Location{}, err
	}

	return location, nil
}
