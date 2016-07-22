package main

import (
    "strconv"
)

const (
    POKEDEX_REGISTRY_URL string = "https://gist.githubusercontent.com/jacobmarshall/4b990717a8d4221586df9f9e68414894/raw/pokedex.json"
)

func LoadPokedex() (map[int]PokedexPokemon, error) {
    unparsed, err := RequestJSON(POKEDEX_REGISTRY_URL)
    if err != nil {
        return nil, err
    }
    parsed := make(map[int]PokedexPokemon)
    for key, value := range unparsed {
        index, err := strconv.Atoi(key)
        if err != nil {
            return nil, err
        }
        parsed[index] = PokedexPokemon{index, value.(string)}
    }
    return parsed, nil
}
