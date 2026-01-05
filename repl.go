package main

import "strings"

func cleanInput(text string) []string {
	split := strings.Split(text, " ")

	var cleanedStrings []string
	for _, splitString := range split {
		if splitString == "" {
			continue
		}

		cleanedStrings = append(cleanedStrings, strings.TrimSpace(strings.ToLower(splitString)))
	}

	return cleanedStrings
}
