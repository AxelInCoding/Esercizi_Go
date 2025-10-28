package main

import "fmt"

func main() {
	/*
	 in questo programma viene dichiarato un num, che viene letto da tastiera
	 se il numero inserito è maggiore di 0 allora viene stampato "positivo"
	 altrimenti se il numero è uguale a 0 allora viene stampato "nullo"
	 altrimenti se entrambe le due condizioni precedenti non sono valide viene stampato "negativo"
	 
	*/

	var num int

	fmt.Print("un int: ")
	fmt.Scan(&num)

	if num > 0 {
		fmt.Println("positivo")
	} else if num == 0 {
		fmt.Println("nullo")
	} else { // < 0
		fmt.Println("negativo")
	}
}

