package main

import (
	"time"
	"fmt"
)

type Pokemon struct {
	ID        int
	ExpiresAt time.Time
	PokedexID int
	Latitude  float64
	Longitude float64
}

func (p *Pokemon) UID() string {
	return fmt.Sprintf("%v-%v-%v-%v", p.ExpiresAt, p.PokedexID, p.Latitude, p.Longitude)
}

func (p *Pokemon) IsVisible() bool {
	return p.ExpiresAt.After(time.Now())
}

func (p *Pokemon) IsInRange(lat, lon float64, distance int) bool {
	return int(DistanceBetween(lat, lon, p.Latitude, p.Longitude)) < distance
}
