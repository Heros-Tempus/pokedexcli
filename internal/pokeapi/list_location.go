package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (RespShallowLocations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	if val, ok := c.cache.Get(url); ok {
		locationsResp := RespShallowLocations{}
		err := json.Unmarshal(val, &locationsResp)
		if err != nil {
			return RespShallowLocations{}, err
		}
		return locationsResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowLocations{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowLocations{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 299 && resp.StatusCode < 500 {
		return RespShallowLocations{}, fmt.Errorf("Client error: %v", resp.StatusCode)
	} else if resp.StatusCode > 499 {
		return RespShallowLocations{}, fmt.Errorf("Server error: %v", resp.StatusCode)
	}

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespShallowLocations{}, err
	}

	locationsResp := RespShallowLocations{}
	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return RespShallowLocations{}, err
	}

	c.cache.Add(url, dat)

	return locationsResp, nil
}

func (c *Client) ExploreLocation(locationName string) (RespDeepLocations, error) {
	url := baseURL + "/location-area/" + locationName

	if val, ok := c.cache.Get(url); ok {
		locationResp := RespDeepLocations{}
		err := json.Unmarshal(val, &locationResp)
		if err != nil {
			return RespDeepLocations{}, err
		}
		return locationResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespDeepLocations{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespDeepLocations{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 299 && resp.StatusCode < 500 {
		return RespDeepLocations{}, fmt.Errorf("Client error: %v, make sure \"%s\" is a valid location name", resp.StatusCode, locationName)
	} else if resp.StatusCode > 499 {
		return RespDeepLocations{}, fmt.Errorf("Server error: %v", resp.StatusCode)
	}

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespDeepLocations{}, err
	}

	locationResp := RespDeepLocations{}
	err = json.Unmarshal(dat, &locationResp)
	if err != nil {
		return RespDeepLocations{}, err
	}

	c.cache.Add(url, dat)

	return locationResp, nil
}
