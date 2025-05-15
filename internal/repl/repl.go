package repl

import "strings"

func CleanInput(text string) []string {
	var result []string

	for _, word := range strings.Fields(text) {
		result = append(result, strings.ToLower(word))
	}
	return result
}
