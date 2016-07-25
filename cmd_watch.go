package main

import (
    "github.com/bugsnag/bugsnag-go"
    "github.com/urfave/cli"
    "log"
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
        cli.StringFlag{
            Name:  "name",
            Usage: "the name of the region for notifying",
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
        cli.BoolFlag{
            Name:  "slack",
            Usage: "send updates to slack (as well as stdout)",
        },
        cli.StringFlag{
            Name:  "slack-webhook-url",
            Usage: "the slack webhook to send the payload to",
        },
        cli.StringFlag{
            Name:  "slack-channel",
            Usage: "override the webhook's default channel",
        },
        cli.BoolFlag{
            Name:  "forever",
            Usage: "keeps retrying even when the API is down",
        },
        cli.IntFlag{
            Name:  "range",
            Usage: "the maximum range to report (in meters)",
            Value: 100,
        },
        cli.BoolFlag{
            Name:  "ignore-common",
            Usage: "prevents notifying you of pokémon which show up litterally everywhere",
        },
    },
}

func WatchAction(c *cli.Context) error {
    config := Config{
        Name: c.String("name"),
        Lat: c.Float64("lat"),
        Lon: c.Float64("lon"),
        Interval: c.Int("interval"),
        ForceInitial: c.Bool("force-initial"),
        Forever: c.BoolT("forever"),
        Range: c.Int("range"),
        IgnoreCommon: c.Bool("ignore-common"),
        Slack: SlackConfig{
            Enable: c.Bool("slack"),
            WebhookURL: c.String("slack-webhook-url"),
            Channel: c.String("slack-channel"),
        },
    }

    handleError := func(err error) {
        if config.Forever {
            time.Sleep(time.Second * time.Duration(5))
        } else {
            bugsnag.Notify(err)
            log.Fatalln(err)
        }
    }

    if config.Interval < 30 {
        config.Interval = 30;
    }

    if config.Lat == 0 && config.Lon == 0 {
        log.Fatalln("Must provide coordinates (pokevision watch --lat=34.00846023931052 --lon=-118.49802017211914)")
    }

    if config.Name == "" {
        log.Fatalln("Must provide a display name for the location (pokevision watch --lat=34.00846023931052 --lon=-118.49802017211914 --name=Beach)")
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
        if count > 1 || config.ForceInitial {
            err = UpdatePokemonInRegion(config.Lat, config.Lon)
            if err != nil {
                handleError(err)
                continue
            }
        }

        current, err := FetchPokemonsInRegion(config.Lat, config.Lon)
        if err != nil {
            handleError(err)
            continue
        }

        new := GetUpdates(pokemon, current).New

        for _, pokemon := range new {
            pokedexPokemon := pokedex.Get(pokemon.PokedexID)
            ignore := false
            if config.IgnoreCommon {
                ignore = pokedexPokemon.IsCommon()
            }
            if !ignore && pokemon.IsVisible() && pokemon.IsInRange(config.Lat, config.Lon, config.Range) {
                meta := GetPokemonMeta(config, pokedex, pokemon)

                OutputToTerminal(meta, config)

                if config.Slack.Enable {
                    go OutputToSlack(meta, config)
                }
            }
        }

        pokemon = current
        time.Sleep(time.Second * time.Duration(config.Interval))
    }
}
