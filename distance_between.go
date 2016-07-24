package main

import (
	"github.com/kellydunn/golang-geo"
	"fmt"
	"strconv"
)

func DistanceBetween(latA, lonA, latB, lonB float64) float64 {
	a := geo.NewPoint(latA, lonA)
	b := geo.NewPoint(latB, lonB)
	return a.GreatCircleDistance(b) * 1000
}

func getDP(num float64, dp int) float64 {
	format := "%." + strconv.Itoa(dp) + "f"
	result, _ := strconv.ParseFloat(fmt.Sprintf(format, num), 64)
	return result
}

func DistanceBetweenDP(latA, lonA, latB, lonB float64, dp int) float64 {
	distance := DistanceBetween(latA, lonA, latB, lonB)
	return getDP(distance, dp)
}

func DistanceBetweenM(latA, lonA, latB, lonB float64) string {
	return fmt.Sprintf("%v", DistanceBetweenDP(latA, lonA, latB, lonB, 0)) + "m"
}