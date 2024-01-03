package main

import (
	"fmt"
	"math/rand"
)

const (
	width  = 80
	height = 15
)

type universe [][]bool

func newUniverse() universe {
	u := make(universe, height)
	for i := range u {
		u[i] = make([]bool, width)
	}

	return u
}

func (u universe) seed() universe {
	liveProhabition := 4
	for i, row := range u {
		for j := range row {
			u[i][j] = (rand.Intn(liveProhabition)+1)%liveProhabition == 0 // ~25% of 'true'
		}
	}

	return u
}

func (u universe) alive(x, y int) bool {
	if x >= 0 && x < width && y >= 0 && y < height {
		return u[y][x]
	}
	return false
}

func (u universe) neighbors(x, y int) int {
	count := 0
	for i := y - 1; i <= y+1; i++ {
		for j := x - 1; j <= x+1; j++ {
			if u.alive(j, i) && !(x == j && y == i) {
				count++
			}
		}
	}
	return count
}

func (u universe) next(x, y int) bool {
	a := u.alive(x, y)
	n := u.neighbors(x, y)
	if a && n < 2 {
		return false
	}
	if a && (n == 2 || n == 3) {
		return true
	}
	if a && n > 3 {
		return false
	}
	if !a && n == 3 {
		return true
	}
	return false
}

func step(a, b universe) {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			b[y][x] = a.next(x, y)
		}
	}
}

func (u universe) show() {
	for i := 0; i < height; i++ {
		var stringRow []string
		for j := 0; j < width; j++ {
			var output string
			if u[i][j] == true {
				output = "*"
			} else {
				output = " "
			}
			stringRow = append(stringRow, fmt.Sprint(output))
		}
		fmt.Println(stringRow)
	}
}

func main() {
	a, b := newUniverse(), newUniverse()
	a.seed()

	for i := 0; i < 300; i++ {
		step(a, b)
		a.show()
		fmt.Println("=================================", i)
		a, b = b, a
	}
}
