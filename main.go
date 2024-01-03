package main

import (
	"fmt"
)

func main() {

	for i, j := 0, 0; i < 15; i++ {
		if j++; j == 3 {
			j = 0
		}
		fmt.Printf("i: %d, j: %d\n", i, j)
	}
}
