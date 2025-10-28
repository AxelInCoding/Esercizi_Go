package main

import (
	"fmt"
	"math/rand"
)

func main() {
	const k = 10;
	var sx1, dx1 int =1,10
	var somma int;

	for i := 1; i <= k; i++ {
		var numeroRandom int = rand.Intn(dx1 - sx1 +1)+ sx1;
		fmt.Print(numeroRandom, " ")
		somma+=numeroRandom;
	}
	fmt.Println();
	fmt.Println("Somma dei numeri estratti: ", somma);
}
