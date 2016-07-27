package main

import (
    notifier "github.com/deckarep/gosx-notifier"
)

func OutputToNotificationCenter(pokemon PokemonMeta, config Config) error {
    pokemonIcon, _ := DownloadIcon(pokemon.PokedexPokemon.Index)

    notification := notifier.Notification{
        Group: "com.pokevision.cli",
        Title: pokemon.Name + " located near " + pokemon.Location,
        Message: "Expires in " + HumanTime(pokemon.ExpiresAt),
        Sound: notifier.Glass,
        Link: pokemon.URL,
        ContentImage: pokemonIcon,
    }

    return notification.Push()
}
