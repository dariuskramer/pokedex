package pokeapi

import (
	"fmt"
	"strings"
)

type Pokemon struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	Weight         int    `json:"weight"`
	Stats          []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Type struct {
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`
}

func (p Pokemon) Formatter() string {
	var builder strings.Builder

	fmt.Fprintf(&builder, "Name: %s\n", p.Name)
	fmt.Fprintf(&builder, "Base experience: %d\n", p.BaseExperience)
	fmt.Fprintf(&builder, "Height: %d\n", p.Height)
	fmt.Fprintf(&builder, "Weight: %d\n", p.Weight)

	builder.WriteString("Stats:\n")
	for _, statInfo := range p.Stats {
		fmt.Fprintf(&builder, "  - %s: %v\n", statInfo.Stat.Name, statInfo.BaseStat)
	}

	builder.WriteString("Types:\n")
	for _, typeInfo := range p.Types {
		fmt.Fprintf(&builder, "  - %s\n", typeInfo.Type.Name)
	}

	return builder.String()
}
