package main

import "fmt"

func main() {
	/* 	 
	In questo programma viene dichiarata una variabile voto intera, viene fatto inserire il voto 
	come input da tastiera. Se il voto è minore oppure maggiore di 100, viene stampato "voto non valido"
	Se la varibile voto è maggiore e uguale a 90, viene stampato A, altrimenti se il voto è maggiore uguale
	a 80 viee stampato B, altrimento se il voto è maggiore uguale di 70 viene stampato C, altrimenti
	se il voto è maggiore uguale di 60 viene stampato D, altrimenti se le condizioni precedenti sono false viene stampato
	F.
	*/

	var voto int
	fmt.Print("voto? ")
	fmt.Scan(&voto)

	if voto < 0 || voto > 100 {
		fmt.Println("voto non valido")
		return //interrompe l'esecuzione del programma
	}

	if voto >= 90 {
		fmt.Println("A")
	} else if voto >= 80 {
		fmt.Println("B")
	} else if voto >= 70 {
		fmt.Println("C")
	} else if voto >= 60 {
		fmt.Println("D")
	} else {
		fmt.Println("F")
	}
}
/* DOMANDE */
/*
- D1. Per quali valori è vera la condizione del primo if? (quello subito dopo la Scan)
- R1. E' vera per i valori minori di 0 oppure per i valori maggiori di 100

- D2. Per quali valori viene eseguita la seguente istruzione?
    fmt.Println("B")
- R2. Viene eseguita l'istruzione della stampa di B solamente se il voto è compreso uguale tra 80 e 90

- D3. Giustificate la vostra risposta alla domanda D2
- R3. Go valuta dell'alto verso il basso, se il voto è >= 90 è vero viene eseguito il Println("A") e nessuno altro bloco viene considerato
  se voto < 90 quindi il primo if è falso e voto >= 80 è vero, allora viene eseguito il Println("B")
  quindi il Pritln("B") non viene eseguito se il voto >= 90 perchè etra nel primo blocco

- D4. Il programma è impostato correttamente o sarebbe stato (più) corretto scrivere invece così?
 } else if voto >= 80 && voto < 90 {
- R4. Il programma è corretto anche così com'è. Non è necessario aggiungere voto < 90 perchè la struttura
if ... else if garantisce già che questo blocco vengo eseguito solo se voto < 90

- D5. Perché?
- R5. Il motivo è la sequenza di if/else if in Go:
  se una condizione è vera, il relativo blocco viene eseguito e tutti gli altri else if successivi vengono ignorati.
  Quindi, quando si arriva a else if voto >= 80, si sa già che voto < 90, perciò non serve ripetere voto < 90 nella condizione:
  sarebbe ridondante
*/

