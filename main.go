package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	initCommands()
	for scanner.Scan() {
		input := scanner.Text()
		cmd, ok := commands[input]

		if !ok {
			fmt.Println("Unknown command")
			continue
		}

		if err := cmd.callback(); err != nil {
			fmt.Println(err)
		}

	}
}
