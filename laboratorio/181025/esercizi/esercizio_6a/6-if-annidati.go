package main

import "fmt"

func main() {
	/*
Il programma chiede all'utente di scegliere una figura geometrica (rettangolo o triangolo) 
e l'operazione da calcolare (perimetro o area). 
In base alla scelta, stampa la formula corrispondente senza eseguire i calcoli.
	*/
	const (
		RETTANGOLO = 1
		TRIANGOLO  = 2
		PERIMETRO  = 1
		AREA       = 2
	)
	var figura, operazione int

	fmt.Println("di che figura si tratta?")
	fmt.Print("rettangolo (1) o triangolo (2)? ")
	fmt.Scan(&figura)
	fmt.Println("cosa vuoi calcolare?")
	fmt.Print("perimetro (1) o area (2)? ")
	fmt.Scan(&operazione)

	if figura == RETTANGOLO {
		if operazione == PERIMETRO {
			fmt.Println("formula = (base + altezza) * 2")
		} else { //operazione == AREA
			fmt.Println("formula = base * altezza")
		}
	} else { //figura == TRIANGOLO
		if operazione == PERIMETRO {
			fmt.Println("formula = lato1 + lato2 + lato3 ")
		} else { //operazione == AREA
			fmt.Println("formula = (base * altezza)/2")
		}
	}
}
/* DOMANDE */
/*
- D1. Quanti diversi messaggi da stampare ci sono?
- R1. Ci sono 4 messaggi diversi da stampare:
      Rettangolo-Perimetro
      Rettangolo-Area
      Triangolo-Perimetro
      Triangolo-Area

- D2. Sarebbe possibile ridurre il numero di if/else? Perché?
- R2. si è possibile ridurre il numero di if/else usando switch case annidati
o combinati per gestire le combinazioni (figura, operazione)
*/

