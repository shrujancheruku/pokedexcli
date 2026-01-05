package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		for scanner.Scan() {
			text := scanner.Text()

			if text == "" {
				break
			}

			cleanedInput := cleanInput(text)

			if len(cleanedInput) == 0 {
				break
			}

			fmt.Printf("Your command was: %s \n", cleanedInput[0])
			break
		}
	}
}
