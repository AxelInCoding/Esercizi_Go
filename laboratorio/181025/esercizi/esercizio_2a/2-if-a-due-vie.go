package main

import "fmt"

func main() {
	/*
	viene dichiarazione una variabile int e viene, e viene fatto inserire un numero da tastiera
	se il numero inserito è divisibile per 2 viene stampato "è pari"
	altrimenti viene stampato "è dispari"
	*/
	
	var n int
	fmt.Print("numero: ")
	fmt.Scan(&n)

	if n%2 == 0 {
		fmt.Println(n, "è pari")
	} else {
		fmt.Println(n, "è dispari")
	}
}

