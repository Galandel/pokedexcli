package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		// Print the prompt
		fmt.Print("Pokedex > ")

		// Wait for user input
		if !scanner.Scan() {
			// Exit the loop if there's an error or EOF
			break
		}

		// Get the user's input and clean it
		input := cleanInput(scanner.Text())

		if len(input) == 0 {
			// fmt.Println()
			continue
		}

		// Print the first word
		fmt.Printf("Your command was: %s\n", input[0])
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
	// return make([]string, 0)
}
