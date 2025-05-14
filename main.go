package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var supportedCommands map[string]cliCommand

func init() {
	supportedCommands = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Display a help message",
			callback:    commandHelp,
		},
	}
}

func main() {
	const Prompt = "Pokedex > "

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(Prompt)
		if !scanner.Scan() {
			log.Fatalf("Scanner error: %v", scanner.Err())
		}
		input := cleanInput(scanner.Text())
		if len(input) == 0 {
			continue
		}

		command, ok := supportedCommands[input[0]]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}

		command.callback()
	}
}

func cleanInput(text string) []string {
	var result []string

	for _, word := range strings.Fields(text) {
		result = append(result, strings.ToLower(word))
	}
	return result
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	usage := `
Welcome to the Pokedex!
Usage:

`
	for _, command := range supportedCommands {
		usage += fmt.Sprintf("%s: %s\n", command.name, command.description)
	}

	fmt.Println(usage)
	return nil
}
