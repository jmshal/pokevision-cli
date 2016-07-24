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
    },
}

func WatchAction(c *cli.Context) error {
    lat := c.Float64("lat")
    lon := c.Float64("lon")
    interval := c.Int("interval")
    forceInitial := c.Bool("force-initial")
    slack := c.Bool("slack")
    slackWebhookUrl := c.String("slack-webhook-url")
    name := c.String("name")
    exitOnError := !c.BoolT("forever")
    slackChannel := c.String("slack-channel")

    handleError := func(err error) {
        if exitOnError {
            bugsnag.Notify(err)
            log.Fatalln(err)
        } else {
            time.Sleep(time.Second * time.Duration(5))
        }
    }

    if interval < 30 {
        interval = 30;
    }

    if lat == 0 && lon == 0 {
        log.Fatalln("Must provide coordinates (pokevision watch --lat=34.00846023931052 --lon=-118.49802017211914)")
    }

    var pokemon []Pokemon
    var slackClient SlackClient

    if slack {
        slackClient = CreateSlackClient(slackWebhookUrl)
    }

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
                handleError(err)
                continue
            }
        }

        current, err := FetchPokemonsInRegion(lat, lon)
        if err != nil {
            handleError(err)
            continue
        }

        new := GetUpdates(pokemon, current).New

        if len(new) > 0 {
            //log.Println("Found new Pokémon...")
        }

        for _, pokemon := range new {
            if pokemon.IsVisible() {
                pokedexPokemon := pokedex[pokemon.PokedexID]
                pokemonName := pokedexPokemon.Name
                pokemonExpiresAtHuman := HumanTime(pokemon.ExpiresAt)
                pokemonLatitude := pokemon.Latitude
                pokemonLongitude := pokemon.Longitude
                pokemonCoords := fmt.Sprintf("%v/%v", pokemonLatitude, pokemonLongitude)
                pokemonDistance := DistanceBetweenM(lat, lon, pokemonLatitude, pokemonLongitude)
                if name != "" {
                    pokemonCoords = fmt.Sprintf("%v (%v)", name, pokemonCoords)
                }

                fmt.Println(fmt.Sprintf("%v (%v) located near %v", pokemonName, pokemonExpiresAtHuman, pokemonCoords))

                if slack {
                    slackClient.SendPayload(SlackPayload{
                        IconURL: pokedexPokemon.Icon(),
                        Markdown: true,
                        Channel: slackChannel,
                        Attachments: []SlackAttachment{
                            {
                                Title: pokemonName,
                                TitleLink: fmt.Sprintf("https://maps.google.com/maps?q=%v,%v&z=19", pokemonLatitude, pokemonLongitude),
                                Fallback: fmt.Sprintf("%v located near %v", pokemonName, name),
                                Fields: []SlackAttachmentField{
                                    {Title: "Expires in", Value: pokemonExpiresAtHuman, Short: true},
                                    {Title: "Location", Value: name, Short: true},
                                    {Title: "Distance", Value: pokemonDistance, Short: true},
                                },
                            },
                        },
                    })
                }
            }
        }

        pokemon = current
        time.Sleep(time.Second * time.Duration(interval))
    }
}
