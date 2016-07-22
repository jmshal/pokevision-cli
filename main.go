package main

import (
    "github.com/urfave/cli"
    "os"
    "time"
)

func main() {
    app := cli.NewApp()
    app.Name = "pokevision"
    app.Usage = "the pokevision cli for watching for new pokémon"
    app.Version = "1.0.0"
    app.Compiled = time.Now()
    app.Authors = []cli.Author{
        cli.Author{
            Name:  "Jacob Marshall",
            Email: "pokemon@jacobmarshall.co",
        },
    }
    app.Copyright = "(c) 2016 Managenet and PokeVision.com"
    app.Commands = []cli.Command{
        WatchCommand,
    }
    app.Run(os.Args)
}
