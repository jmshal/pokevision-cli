package main

import (
    "github.com/bugsnag/bugsnag-go"
    "github.com/urfave/cli"
    "log"
    "fmt"
    "time"
)

var WatchCommand = cli.Command{
    Name:   "watch",
    Usage:  "watches for new pokémon in a region",
    Action: WatchAction,
    Flags:  []cli.Flag {
        cli.Float64Flag{
            Name:  "lon",
            Usage: "the region's longitude coordinate",
        },
        cli.Float64Flag{
            Name:  "lat",
            Usage: "the region's latitude coordinate",
        },
        cli.IntFlag{
            Name:  "interval",
            Value: 30,
            Usage: "the interval (in seconds) to check for new pokémon",
        },
        cli.BoolFlag{
            Name:  "force-initial",
            Usage: "forces the initial region fetch to check for new pokémon (slower initially)",
        },
    },
}

func WatchAction(c *cli.Context) error {
    lat := c.Float64("lat")
    lon := c.Float64("lon")
    interval := c.Int("interval")
    forceInitial := c.Bool("force-initial")

    if interval < 30 {
        interval = 30;
    }

    if lat == 0 && lon == 0 {
        log.Fatalln("Must provide coordinates (pokevision watch --lat=34.00846023931052 --lon=-118.49802017211914)")
    }

    var pokemon []Pokemon

    pokedex, err := LoadPokedex()
    if err != nil {
        bugsnag.Notify(err)
        log.Fatalln(err)
    }

    count := 0
    for {
        count++

        // Let the first request be the quickest, updating pokemon regions can be slow
        if count > 1 || forceInitial {
            err = UpdatePokemonInRegion(lat, lon)
            if err != nil {
                bugsnag.Notify(err)
                log.Fatalln(err)
            }
        }

        current, err := FetchPokemonsInRegion(lat, lon)
        if err != nil {
            bugsnag.Notify(err)
            log.Fatalln(err)
        }

        new := GetUpdates(pokemon, current).New

        if len(new) > 0 {
            log.Println("Found new Pokémon...")
        }

        for _, pokemon := range new {
            if pokemon.IsVisible() {
                pokedexPokemon := pokedex[pokemon.PokedexID]
                log.Println(fmt.Sprintf("%v (%v) at %v/%v", pokedexPokemon.Name, HumanTime(pokemon.ExpiresAt), pokemon.Latitude, pokemon.Longitude))
            }
        }

        pokemon = current
        time.Sleep(time.Second * time.Duration(interval))
    }
}
