package main

import (
    "strconv"
    "fmt"
    "encoding/json"
)

const (
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

func LoadPokedex(locale string) (Pokedex, error) {
    pokedex := Pokedex{}
    file, err := Asset("pokedex." + locale + ".json")
    if err != nil {
        return pokedex, err
    }
    var unparsed map[string]interface{}
    err = json.Unmarshal(file, &unparsed)
    if err != nil {
        return pokedex, err
    }
    pokedex.index = make(map[int]PokedexPokemon)
    for key, value := range unparsed {
        index, err := strconv.Atoi(key)
        if err != nil {
            return Pokedex{}, err
        }
        pokedex.index[index] = PokedexPokemon{index, value.(string)}
    }
    return pokedex, nil
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
