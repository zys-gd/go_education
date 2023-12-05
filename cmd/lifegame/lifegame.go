package main

import (
	"fmt"
)

const (
	width  = 80
	height = 15
)

type Universe [][]bool

func newUniverse() Universe {
	u := make(Universe, height)
	for i := range u {
		u[i] = make([]bool, width)
	}

	return u
}

func show(u Universe) {

	for _, row := range u {
		var stringRow []string
		for _, val := range row {
			var output string
			if val == true {
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
	show(newUniverse())
}
