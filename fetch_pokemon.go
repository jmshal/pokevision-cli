package main

import (
    "fmt"
    "time"
    "strconv"
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
        pokemonId, _ := strconv.Atoi(rawPokemon["pokemonId"].(string))
        pokemons[index] = Pokemon{
            ID: rawPokemon["id"].(string),
            ExpiresAt: time.Unix(int64(rawPokemon["expiration_time"].(float64)), 0),
            PokedexID: pokemonId,
            Latitude: rawPokemon["latitude"].(float64),
            Longitude: rawPokemon["longitude"].(float64),
            IsAlive: rawPokemon["is_alive"].(bool),
        }
    }
    return pokemons, nil
}
