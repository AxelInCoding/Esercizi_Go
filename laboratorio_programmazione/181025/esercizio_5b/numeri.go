package main

import "fmt"

/*
Per il programma numeri.go vogliamo stampare messaggi diversi a seconda delle proprietà del numero inserito:

- "positivo divisibile per 10" se il numero è sia positivo che divisibile per 10
- "divisibile per 10" se il numero è divisibile per 10 ma non positivo
- "positivo" se il numero è positivo ma non divisibile per 10
- niente altrimenti

Per gestire correttamente questi casi sovrapposti, serve un if a più vie (if ... else if ... else) per evitare che più messaggi vengano stampati per lo stesso input.  

- If a più vie: 1 (per distinguere i 3 casi più il caso "nessuna stampa")  
- If a una via: opzionale, ma non necessario in questa soluzione.
*/

func main() {
	var num int
	fmt.Print("Inserisci un numero: ")
	fmt.Scan(&num)

	if num > 0 && num%10 == 0 {
		fmt.Println("positivo divisibile per 10")
	} else if num%10 == 0 {
		fmt.Println("divisibile per 10")
	} else if num > 0 {
		fmt.Println("positivo")
	}
}

