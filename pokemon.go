package main

import "time"

type Pokemon struct {
    ID        string         `json:"id"`
    ExpiresAt time.Time      `json:"expiration_time"`
    PokedexID int            `json:"pokemonId"`
    Latitude  float64        `json:"latitude"`
    Longitude float64        `json:"longitude"`
    IsAlive   bool           `json:"is_alive"`
}

func (p *Pokemon) IsVisible() bool {
    return p.IsAlive && p.ExpiresAt.After(time.Now())
}
