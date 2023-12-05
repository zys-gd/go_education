package main

import (
	"fmt"
	"math/rand"
)

var move interface {
	move(m []string) string
}

var feed interface {
	feed() string
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

func (a animal) move(m []string) string {
	randomIndex := rand.Intn(len(m))
	return fmt.Sprintf("The animal '%v' is doing now: %v", a.name, m[randomIndex])
}

func main() {
	// Имплементируйте цикл день/ночь, симулятор должен запускаться для трех 24-часовых марсианских дней (72 часа).
	// Все животные должны спать от заката до рассвета.
	// Для каждого часа дня выберите случайное животное, что будет осуществлять какое-то случайное действие — передвигаться или есть.
	// Для каждого действия должно выводиться описание того, что именно сделал зверь.
	a := animal{"dog", activeTime(diurnal)}

	movementList := []string{
		"go left",
		"go right",
		"sit down",
		"play",
		"looking for a food",
	}

	fmt.Println(a.move(movementList))
}
