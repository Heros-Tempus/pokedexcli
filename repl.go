package main

import "strings"

func cleanInput(text string) []string {
	//takes a string as input and returns a slice of lowercase strings split on whitespace
	lowercase := strings.ToLower(text)
	return strings.Fields(lowercase)
}
