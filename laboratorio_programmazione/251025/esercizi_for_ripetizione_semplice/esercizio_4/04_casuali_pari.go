package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var dx1 int = 10
	var contatorePari int;

	for i := 1; i <= 10; i++ {
		var numeroRandom int = rand.Intn(dx1+1);
		fmt.Print(numeroRandom, " ")
    if numeroRandom % 2 == 0 {
      contatorePari++;
		}
	}
	fmt.Println();
	fmt.Print("Numeri pari generati: ", contatorePari);
}
