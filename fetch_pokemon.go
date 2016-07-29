package main

import (
    "fmt"
    "time"
)

const (
    GET_MAP_DATA_URL = "https://pokevision.com/map/data/%v/%v"
)

func FetchPokemonsInRegion(lat, lon float64) ([]Pokemon, error) {
    data, err := RequestAPI(fmt.Sprintf(GET_MAP_DATA_URL, lat, lon))
    if err != nil {
        return nil, err
    }
    rawPokemons := data["pokemon"].([]interface{})
    pokemons := make([]Pokemon, len(rawPokemons))
    for index, unusableRawPokemon := range rawPokemons {
        rawPokemon := unusableRawPokemon.(map[string]interface{})
        pokemonId := int(rawPokemon["pokemonId"].(float64))
        pokemons[index] = Pokemon{
            ID: int(rawPokemon["id"].(float64)),
            ExpiresAt: time.Unix(int64(rawPokemon["expiration_time"].(float64)), 0),
            PokedexID: pokemonId,
            Latitude: rawPokemon["latitude"].(float64),
            Longitude: rawPokemon["longitude"].(float64),
        }
    }
    return pokemons, nil
}
