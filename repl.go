package main

import "strings"

// func replLoop() {
// }

func CleanInput(text string) []string {
	lower := strings.ToLower(text)
	words := strings.Fields(lower)
	return words
}
