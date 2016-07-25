package main

import (
    "strconv"
    "fmt"
)

const (
    POKEDEX_REGISTRY_URL string = "https://gist.githubusercontent.com/jacobmarshall/4b990717a8d4221586df9f9e68414894/raw/pokedex.json"
    POKEDEX_POKEMON_ICON_URL string = "https://ugc.pokevision.com/images/pokemon/%v.png"
)

var (
    // Rattata
    // Pidgey
    // Weedle
    // Caterpie
    // Zubat
    COMMON_POKEMON []int = []int{19, 16, 13, 10, 41}
)

type PokedexPokemon struct {
    Index int
    Name  string
}

type Pokedex struct {
    index map[int]PokedexPokemon
}

func LoadPokedex() (Pokedex, error) {
    unparsed, err := RequestJSON(POKEDEX_REGISTRY_URL)
    if err != nil {
        return Pokedex{}, err
    }
    parsed := make(map[int]PokedexPokemon)
    for key, value := range unparsed {
        index, err := strconv.Atoi(key)
        if err != nil {
            return Pokedex{}, err
        }
        parsed[index] = PokedexPokemon{index, value.(string)}
    }
    return Pokedex{parsed}, nil
}

func (dex *Pokedex) Get(index int) PokedexPokemon {
    return dex.index[index]
}

func (p *PokedexPokemon) Icon() string {
    return fmt.Sprintf(POKEDEX_POKEMON_ICON_URL, p.Index)
}

func (p *PokedexPokemon) IsCommon() bool {
    for _, index := range COMMON_POKEMON {
        if p.Index == index {
            return true
        }
    }
    return false
}