package main

import "fmt"

    /*
      Il programma legge una sequenza di numeri interi da input.
A ogni passo, combina le ultime due cifre dell’ultimo numero letto
con le ultime due cifre del numero precedente per formare un nuovo valore (c).
Il ciclo termina quando questo valore è uguale a 100.
Stampa il valore calcolato a ogni iterazione.
    */

func main() {
	var combinato, numero, precedente int
	fmt.Scan(&numero)
	for combinato != 100 {
		precedente= numero
		fmt.Scan(&numero)
		combinato = (precedente%100)*100 + numero%100
		fmt.Println(combinato)
	}
}

