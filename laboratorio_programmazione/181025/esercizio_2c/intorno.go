package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y float64

	fmt.Scan(&x, &y)

	// Costante di soglia
	const EPSILON = 1e-5

	// Calcolo della distanza dall'origine
	distanza := math.Sqrt(x*x + y*y)

	// Verifica se il punto è molto vicino all'origine
	if distanza < EPSILON {
		fmt.Println("molto vicino all'origine")
	} else {
		fmt.Println("non vicino all'origine")
	}
}

