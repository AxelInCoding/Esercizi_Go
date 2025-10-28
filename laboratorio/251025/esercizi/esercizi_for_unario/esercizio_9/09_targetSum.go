package main

import (
	"fmt"
	"math/rand"
)

func main() {
	const TARGET = 50
	const SX = 1
	const DX = 10

	var somma int

	for somma <= TARGET {
		numeroRandom := rand.Intn(DX-SX+1) + SX
		fmt.Print(numeroRandom, " ")
		somma += numeroRandom
	}

	fmt.Println("\nsomma dei numeri:", somma)
}

