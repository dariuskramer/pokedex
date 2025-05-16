package commands

import (
	"fmt"
	"os"
)

func CommandExit(config *CommandConfig, args ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
