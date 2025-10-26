package main

import "fmt"

func main() {
	/* 
	Questo programma richiede un numero intero da input da tastiera, se il numero è divisibile per 3
	allora viene stampato "Fizz" se invece il numero è divisibile per 5 viene stampato "Buzz"
	*/

	var num int
	fmt.Print("numero? ")
	fmt.Scan(&num)

	if num%3 == 0 {
		fmt.Println("Fizz")
	}
	if num%5 == 0 {
		fmt.Println("Buzz")
	}
}
/* DOMANDE */
/*
- D1. Se al posto del secondo if ci fosse un else if (legato al primo if), il programma si comporterebbe nello stesso modo? Dareste le stesse specifiche?
- R1. No il programma non si comporterebbe nello stesso modo, perciò non darei le stesse spefiche

- D2. Perché? Spiegate la vostra risposta alla domanda D1
- R2. Attualmente ci sono due if separati, se sostituissimo il secondo if con else if 
			se il numero è divisibile per 3, il blocco else if non viene più valutato, quindi per numeri come il 15 stamperebbe solo
			"Fizz" e mai "Buzz"

- D3. Cosa stampa il programma per ciascuno dei seguenti input: 8, 9 10, 15?
- R3  
   8: nessuno dei due if è vero, quindi non stampa nulla
   9: divisibile per 3 quindi stampa "Fizz"
  10: divisibile per 5 quindi stampa "Buzz"
  15: divisibile per 3 e per 5 stampa "Fizz" e "Buzz" (Su due righe separate)
 

*/

