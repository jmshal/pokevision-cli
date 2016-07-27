package main

import (
	"time"
	"fmt"
)

type PokemonMeta struct {
	Name string
	Icon string
	ExpiresAt time.Time
	Latitude, Longitude float64
	Distance int
	Location string
	URL string
	Pokemon Pokemon
	PokedexPokemon PokedexPokemon
}

func GetPokemonMeta(config Config, pokedex Pokedex, pokemon Pokemon) PokemonMeta {
	pokedexPokemon := pokedex.Get(pokemon.PokedexID)
	return PokemonMeta{
		Name: pokedexPokemon.Name,
		Icon: pokedexPokemon.Icon(),
		ExpiresAt: pokemon.ExpiresAt,
		Latitude: pokemon.Latitude,
		Longitude: pokemon.Longitude,
		Distance: int(DistanceBetweenDP(config.Lat, config.Lon, pokemon.Latitude, pokemon.Longitude, 0)),
		Location: config.Name,
		URL: fmt.Sprintf("https://maps.google.com/maps?q=%v,%v&z=19", pokemon.Latitude, pokemon.Longitude),
		Pokemon: pokemon,
		PokedexPokemon: pokedexPokemon,
	}
}
