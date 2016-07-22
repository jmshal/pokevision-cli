package main

import (
    "log"
    "fmt"
    "time"
)

func main() {
    const lat, lon float64 = 34.00846023931052, -118.49802017211914
    var pokemon []Pokemon

    pokedex, err := LoadPokedex()
    if err != nil {
        log.Fatalln(err)
    }

    count := 0
    for {
        count++

        // Let the first request be the quickest, updating pokemon regions can be slow
        if count > 1 {
            err = UpdatePokemonInRegion(lat, lon)
            if err != nil {
                log.Fatalln(err)
            }
        }

        current, err := FetchPokemonsInRegion(lat, lon)
        if err != nil {
            log.Fatalln(err)
        }

        new := GetUpdates(pokemon, current).New

        if len(new) > 0 {
            log.Println("Found new Pokémon...")
        }

        for _, pokemon := range new {
            if pokemon.IsVisible() {
                pokedexPokemon := pokedex[pokemon.PokedexID]
                log.Println(fmt.Sprintf("%v (%v) at %v/%v", pokedexPokemon.Name, HumanTime(pokemon.ExpiresAt), pokemon.Latitude, pokemon.Longitude))
            }
        }

        pokemon = current
        time.Sleep(time.Second * 20)
    }
}
