package main

import (
    "github.com/dustin/go-humanize"
    "time"
    "strings"
)

func HumanTime(time time.Time) string {
    original := humanize.Time(time)
    parts := strings.Split(original, " ")
    return parts[0] + " " + parts[1]
}
