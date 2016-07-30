package main

import (
	"github.com/fatih/color"
	"fmt"
	"log"
	"text/template"
	"bytes"
)

func OutputToTerminal(pokemon PokemonMeta, config Config) {
	if config.Format != "" {
		formatTemplate := template.New("format")
		_, err := formatTemplate.Parse(config.Format)
		if err != nil {
			log.Fatalln(err)
		}

		var output bytes.Buffer
		err = formatTemplate.Execute(&output, pokemon)
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println(output.String())
	} else {
		// Ignore the ugliness :)
		color.New(color.Bold).Print(pokemon.Name)
		color.New().Print(" (")
		color.New(color.FgBlue, color.Underline).Printf("%v", pokemon.URL)
		color.New().Print(")")
		fmt.Println()
		color.New().Printf("%vm away from %v\n", pokemon.Distance, pokemon.Location)
		color.New().Printf("Expires in %v\n", HumanTime(pokemon.ExpiresAt))
		fmt.Println()
	}
}
