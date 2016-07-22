package main

import "time"

type Pokemon struct {
    ID        string
    ExpiresAt time.Time
    PokedexID int
    Latitude  float64
    Longitude float64
    IsAlive   bool
}

func (p *Pokemon) IsVisible() bool {
    return p.IsAlive && p.ExpiresAt.After(time.Now())
}
