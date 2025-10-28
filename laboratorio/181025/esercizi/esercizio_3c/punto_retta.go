package main

import "fmt"
import "math"

func main() {
	var x1, y1, m, q float64

	// Lettura dei dati (senza prompt)
	fmt.Scan(&x1, &y1, &m, &q)

	// Costante di tolleranza per i confronti con float64
	const EPSILON = 1e-6

	// Calcolo del valore della retta nel punto x1
	y_retta := m*x1 + q

	// Differenza tra la y del punto e quella della retta
	diff := y1 - y_retta

	// Verifica della posizione del punto rispetto alla retta
	if math.Abs(diff) < EPSILON {
		fmt.Println("sulla retta")
	} else if diff > 0 {
		fmt.Println("sopra")
	} else {
		fmt.Println("sotto")
	}
}

