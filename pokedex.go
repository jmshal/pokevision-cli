package main

import (
    "strconv"
    "fmt"
)

const (
    POKEDEX_REGISTRY_URL string = "https://gist.githubusercontent.com/jacobmarshall/4b990717a8d4221586df9f9e68414894/raw/pokedex.json"
    POKEDEX_POKEMON_ICON_URL string = "https://ugc.pokevision.com/images/pokemon/%v.png"
)

type PokedexPokemon struct {
    Index int
    Name  string
}

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

func (p *PokedexPokemon) Icon() string {
    return fmt.Sprintf(POKEDEX_POKEMON_ICON_URL, p.Index)
}