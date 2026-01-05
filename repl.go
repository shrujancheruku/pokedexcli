package main

import (
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var commands map[string]cliCommand

func initCommands() {
	commands = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Displays a help message",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Welcome to the Pokedex!",
			callback:    helpCallback,
		},
	}
}

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

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)

	return nil
}

func helpCallback() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	for _, command := range commands {
		fmt.Println(command.name + ": " + command.description)
	}

	return nil
}
