package main

import (
    "os/exec"
    "strconv"
)

func OutputToNotificationCenter(pokemon PokemonMeta, config Config) error {
    pokemonIcon, _ := DownloadIcon(pokemon.PokedexPokemon.Index)
    exec.Command("notify-send", "-i", pokemonIcon, pokemon.Name + " - " + strconv.Itoa(pokemon.Distance) + "m from " + pokemon.Location, "Expires in " + HumanTime(pokemon.ExpiresAt)).Run()
    return nil
}
