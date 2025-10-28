package main

import "fmt"

func main() {
	/*
    In questo programma viene fatto inserire un numero intero tra 1 e 4,
    poi viene fatto un for ternario il quali itera in base al numero inserito * la BASE (una constante = 8)
    stampando ad ogni iterazione delle PIPE con \t.
    Successivamente viene fatto un altro for ternario come quello precedente, 
    questa volta però viene stampato il numero i-esimo/BASE, il resto di questa divisione messi tutti distanti di \t.
	*/

	const BASE = 8
	var n int

	fmt.Print("un int tra 1 e 4: ")
	fmt.Scan(&n)

	for i := 0; i < n*BASE; i++ { // i < n*BASE
		fmt.Print("|\t")
	}
	fmt.Println()

	for i := 0; i < n*BASE; i++ {
		fmt.Print(i/BASE, i%BASE, "\t")
	}
	fmt.Println()
}


