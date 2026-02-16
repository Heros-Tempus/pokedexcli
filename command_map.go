package main

import (
	"fmt"
)

func getLocations(cfg *config, args ...string) error {
	locationsResp, err := cfg.client.ListLocations(cfg.nextLocationsURL)
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
