package main

import (
    "github.com/bugsnag/bugsnag-go"
    "github.com/urfave/cli"
    "os"
    "time"
)

const VERSION = "1.0.3"

func main() {
    bugsnag.Configure(bugsnag.Configuration{
        // Hardcoded for now, but what the hay...
        APIKey: "e3e34aa93a77392fd9daca3c266fc123",
        AppVersion: VERSION,
    })
    defer bugsnag.AutoNotify()

    app := cli.NewApp()
    app.Name = "pokevision"
    app.Usage = "the (unofficial) pokevision cli"
    app.Version = VERSION
    app.Compiled = time.Now()
    app.Authors = []cli.Author{
        cli.Author{
            Name:  "Jacob Marshall",
            Email: "pokemon@jacobmarshall.co",
        },
    }
    app.Commands = []cli.Command{
        WatchCommand,
    }
    app.Run(os.Args)
}
