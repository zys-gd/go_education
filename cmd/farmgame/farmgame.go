package main

import (
	"fmt"
	"math/rand"
)

type move interface {
	move(m []string) string
}

type feed interface {
	feed() string
}

type sleep interface {
	feed() string
}

type animalLifecycle interface {
	move
	feed
	sleep
}

const (
	diurnal   = "diurnal"
	nocturnal = "nocturnal"
)

type activeTime string

type animal struct {
	name       string
	activeTime activeTime
}

func (a animal) feed() string {
	return "eating"
}

func (a animal) sleep() string {
	return fmt.Sprintf("The animal '%v' is sleeping now", a.name)
}

func (a animal) move(m []string) string {
	return fmt.Sprintf("The animal '%v' is doing now: %v", a.name, m[rand.Intn(len(m))])
}

func animalActivity(a animal, isAwake bool, m []string) string {
	if isAwake {
		if rand.Intn(2) == 0 {
			return a.move(m)
		} else {
			return a.feed()
		}
	}

	return a.sleep()
}

func main() {
	a := animal{"dog", activeTime(diurnal)}

	movementList := []string{
		"go left",
		"go right",
		"sit down",
		"play",
		"looking for a food",
	}
	sunrize := 7
	dawn := 19
	hoursInDay := 24
	days := 3
	for i, c := 0, 0; i <= hoursInDay; i++ {
		if i == hoursInDay {
			i = 0
			c++
		}
		if c == days {
			break
		}
		isDay := i >= sunrize && i <= dawn
		isAwake := (a.activeTime == diurnal && isDay == true) || (a.activeTime == nocturnal && isDay == false)

		fmt.Printf("Day: %d; Time of a day: %dh; ", c+1, i)
		fmt.Println(animalActivity(a, isAwake, movementList))
	}
}
