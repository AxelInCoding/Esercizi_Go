package main

import "fmt"

func main() {
	var g1, g2, g3 int

	fmt.Print("Guasto componente 1: ")
	fmt.Scan(&g1)
	fmt.Print("Guasto componente 2: ")
	fmt.Scan(&g2)
	fmt.Print("Guasto componente 3: ")
	fmt.Scan(&g3)

	// Controllo componente 1
	if g1 == 1 {
		fmt.Println("caricamento acqua")
	} else if g1 == 2 {
		fmt.Println("scarico acqua")
	} else if g1 == 3 {
		fmt.Println("sistema di riscaldamento")
	}

	// Controllo componente 2
	if g2 == 1 {
		fmt.Println("caricamento acqua")
	} else if g2 == 2 {
		fmt.Println("scarico acqua")
	} else if g2 == 3 {
		fmt.Println("sistema di riscaldamento")
	}

	// Controllo componente 3
	if g3 == 1 {
		fmt.Println("caricamento acqua")
	} else if g3 == 2 {
		fmt.Println("scarico acqua")
	} else if g3 == 3 {
		fmt.Println("sistema di riscaldamento")
	}
}

