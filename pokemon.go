package main

import "time"

type Pokemon struct {
    ID        int
    ExpiresAt time.Time
    PokedexID int
    Latitude  float64
    Longitude float64
    IsAlive   bool
}

func (p *Pokemon) IsVisible() bool {
    return p.IsAlive && p.ExpiresAt.After(time.Now())
}

func (p *Pokemon) IsInRange(lat, lon float64, distance int) bool {
    return int(DistanceBetween(lat, lon, p.Latitude, p.Longitude)) < distance
}