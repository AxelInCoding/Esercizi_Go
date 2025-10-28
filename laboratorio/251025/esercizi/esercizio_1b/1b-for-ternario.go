package main

import "fmt"

func main() {

	var first, last, step int = 1, 10, 1

	/* Qui sotto trovate quattro commenti (a., b., c., d.) e poi quattro cicli for.
	Per ciascun ciclo for scegliete il commento corrispondente a quello che fa il ciclo for e spostatelo (taglia e incolla) al posto della riga "//...". */

	//	
	//	
	//	
	//	

	// ciclo 1
	// d. stampa una riga orizzontale di asterischi
	for i := first; i <= last; i += step {
		fmt.Print("*")
	}

	fmt.Println()

	// ciclo 2
	// b. stampa una riga vericale di asterischi
	for i := first; i <= last; i += step {
		fmt.Println("*")
	}

	// ciclo 3
	// c. stampa i numeri da 1 a 10, separati da tab
	for i := first; i <= last; i += step {
		fmt.Print(i, "\t")
	}

	fmt.Println()

	// ciclo 4
	// a. produce 10 stampe del valore di first (1), separate da un tab
	for i := first; i <= last; i += step {
		fmt.Print(first, "\t")
	}

	fmt.Println()
}

