package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var sx1, dx1 int = 0, 10

	fmt.Println("Numeri casuali tra", sx1, "e", dx1)
	for i := 1; i <= 10; i++ {
		fmt.Print(rand.Intn(dx1+1), " ")
	}
	fmt.Println()
}
