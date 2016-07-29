package main


import (
	"os/exec"
)

func OutputToNotificationCenter(pokemon PokemonMeta, config Config) error {
    pokemonIcon, _ := DownloadIcon(pokemon.PokedexPokemon.Index)

    exec.Command("notify-send", "-i", pokemonIcon, pokemon.Name + " located near " + pokemon.Location, "Expires in " + HumanTime(pokemon.ExpiresAt)).Run()
    return nil
}
