package main

import (
	"fmt"
	"math"
)

type coordinate struct {
	d, m, s float64
	h       rune
}

type location struct {
	lat, long float64
}

type world struct {
	radius float64
}

func (c coordinate) decimal() float64 {
	sign := 1.0
	switch c.h {
	case 'S', 'W', 's', 'w':
		sign = -1
	}
	return sign * (c.d + c.m/60 + c.s/3600)
}

func (w world) distance(p1, p2 location) float64 {
	s1, c1 := math.Sincos(rad(p1.lat))
	s2, c2 := math.Sincos(rad(p2.lat))
	clong := math.Cos(rad(p1.long - p2.long))
	return w.radius * math.Acos(s1*s2+c1*c2*clong)
}

func newLocation(lat, long coordinate) location {
	return location{lat.decimal(), long.decimal()}
}

func rad(deg float64) float64 {
	return deg * math.Pi / 180
}

func main() {
	var mars = world{radius: 3389.5}
	list := map[string]location{
		"Spirit":      {coordinate{14, 34, 6.2, 'S'}.decimal(), coordinate{175, 28, 21.5, 'E'}.decimal()},
		"Opportunity": {coordinate{1, 56, 46.3, 'S'}.decimal(), coordinate{354, 28, 24.2, 'E'}.decimal()},
		"Curiosity":   {coordinate{4, 35, 22.2, 'S'}.decimal(), coordinate{137, 26, 30.1, 'E'}.decimal()},
		"InSight":     {coordinate{4, 30, 0.0, 'N'}.decimal(), coordinate{135, 54, 0, 'E'}.decimal()},
	}
	var dist, closestDist, farthestDist float64
	var closestSpot, farthestSpot string

	for rover, loc := range list {
		for rover2, loc2 := range list {
			if rover == rover2 {
				continue
			}
			dist = mars.distance(loc, loc2)
			fmt.Printf("%s - %s: %v \n", rover, rover2, dist)
			if closestDist == 0 || closestDist > dist {
				closestDist = dist
				closestSpot = fmt.Sprintf("%s - %s", rover, rover2)
			}
			if farthestDist == 0 || farthestDist < dist {
				farthestDist = dist
				farthestSpot = fmt.Sprintf("%s - %s", rover, rover2)
			}
			delete(list, rover)
		}
	}

	fmt.Printf("\nClosest: %v: %v \n", closestSpot, closestDist)
	fmt.Printf("Farthest: %v: %v \n", farthestSpot, farthestDist)
}
