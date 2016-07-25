package main

import (
	"fmt"
	"github.com/fatih/color"
)

func OutputToTerminal(pokemon PokemonMeta, config Config) {
	// Ignore the ugliness :)
	//color.New(color.BgWhite, color.FgBlue, color.Bold).Print(pokemon.Name)
	color.New(color.Bold).Print(pokemon.Name)
	color.New().Print(" (")
	color.New(color.FgBlue, color.Underline).Printf("%v", pokemon.URL)
	color.New().Print(")")
	fmt.Println()
	color.New().Printf("%vm away from %v\n", pokemon.Distance, pokemon.Location)
	color.New().Printf("Expires in %v\n", HumanTime(pokemon.ExpiresAt))
	fmt.Println()
}