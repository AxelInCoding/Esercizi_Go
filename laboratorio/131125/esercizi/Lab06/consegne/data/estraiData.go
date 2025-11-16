/*
1. Analizzatelo e rispondete alle seguenti domande

Di che tipo è la variabile mesi?
-E' un array di stringhe
A cosa serve la variabile mesi?
-A determinare il mese corrente
Cosa fa la funzione num2mese?
-Ritorna il numero intero inserito in input col nome del mese corrispondente
*/

/*Codice del programma estraiData*/
package main

import (
   "fmt"
   "strconv"
   "strings"
)

func main() {
   var dataGMA string
   fmt.Print("data nel formato giorno/mese/anno: ")
   fmt.Scan(&dataGMA)
   g, m, a := stringGMA2GMA(dataGMA)
   fmt.Println("giorno", g)
   fmt.Println("mese", num2mese(m))
   fmt.Println("anno", a)
}

func stringGMA2GMA(data string) (int, int, int) {
	// Suddivido la stringa nei tre componenti
	parts := strings.Split(data, "/")
	if len(parts) != 3 {
		// Gestione minimale di errore: ritorna zeri
		return 0, 0, 0
	}

	// Converte ciascun elemento da stringa a intero
	giorno, _ := strconv.Atoi(parts[0])
	mese, _ := strconv.Atoi(parts[1])
	anno, _ := strconv.Atoi(parts[2])

	return giorno, mese, anno
}

func num2mese(m int) string {
   var mesi = [13]string{"", "gen", "feb", "mar", "apr", "mag", "giu", "lug", "ago", "set", "ott", "nov", "dic"}
   return mesi[m]
}

