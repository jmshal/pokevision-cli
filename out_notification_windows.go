package main

import (
    toast "github.com/jacobmarshall/go-toast"
    "fmt"
)

func OutputToNotificationCenter(pokemon PokemonMeta, config Config) error {
    pokemonIcon, _ := DownloadIcon(pokemon.PokedexPokemon.Index)
    mapURI := fmt.Sprintf("bingmaps:?collection=point.%v_%v_%v", pokemon.Latitude, pokemon.Longitude, pokemon.Name)

    notification := toast.Notification{
        AppID: "PokeVision", // Shows up in the action center (lack of accent is due to encoding issues)
        Title: pokemon.Name + " located near " + pokemon.Location,
        Message: "Expires in " + HumanTime(pokemon.ExpiresAt),
        Icon: pokemonIcon,
        Actions: []toast.Action{
            {"protocol", "View on map", mapURI},
        },
    }

    return notification.Push()
}
