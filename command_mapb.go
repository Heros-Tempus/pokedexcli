package main

import (
	"fmt"
)

func getPreviousLocations(cfg *config, args ...string) error {
	if cfg.prevLocationsURL == nil {
		return fmt.Errorf("you're on the first page")
	}
	locationsResp, err := cfg.client.ListLocations(cfg.prevLocationsURL)
	if err != nil {
		return err
	}
	cfg.nextLocationsURL = locationsResp.Next
	cfg.prevLocationsURL = locationsResp.Previous

	for _, r := range locationsResp.Results {
		fmt.Println(r.Name)
	}
	return nil
}
