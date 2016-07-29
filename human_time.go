package main

import (
	"github.com/dustin/go-humanize"
	"time"
	"strings"
	"math"
)

func HumanTime(time time.Time) string {
	original := humanize.Time(time)
	parts := strings.Split(original, " ")
	length := int(math.Min(float64(len(parts)), 2)) // now / x minutes
	parts = parts[:length]
	return strings.Join(parts, " ")
}
