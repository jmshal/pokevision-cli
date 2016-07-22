package main

import (
    "github.com/urfave/cli"
    "os"
    "time"
)

func main() {
    app := cli.NewApp()
    app.Name = "pokevision"
    app.Usage = "the (unofficial) pokevision cli"
    app.Version = "1.0.1"
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
