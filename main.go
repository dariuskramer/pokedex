package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/dariuskramer/pokedex/internal/commands"
	"github.com/dariuskramer/pokedex/internal/pokecache"
	"github.com/dariuskramer/pokedex/internal/repl"
)

func main() {
	const Prompt = "Pokedex > "
	const cacheDuration = 5 * time.Second
	config := &commands.CommandConfig{
		Cache: pokecache.NewCache(cacheDuration),
	}
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(Prompt)
		if !scanner.Scan() {
			log.Fatalf("Scanner error: %v", scanner.Err())
		}
		input := repl.CleanInput(scanner.Text())
		if len(input) == 0 {
			continue
		}

		command, ok := commands.SupportedCommands[input[0]]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}

		var args []string
		if len(input) > 1 {
			args = input[1:]
		}

		err := command.Callback(config, args)
		if err != nil {
			fmt.Println(err)
		}
	}
}
