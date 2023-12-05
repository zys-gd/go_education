package main

import (
	"fmt"
	"math"
)

func main() {
	temperatures := []float64{
		-28.0, 32.0, -31.0, -29.0, -23.0, -29.0, -28.0, -33.0,
	}

	groups := make(map[float64][]float64) // Карта с ключами float64 и значениями []float64

	for _, t := range temperatures {
		g := math.Trunc(t/10) * 10 // Округляет температуры вниз -20, -30 и так далее
		groups[g] = append(groups[g], t)
	}

	fmt.Print(groups)
}
