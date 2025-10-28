/*Risposta alla domanda:
Sì, questo problema può essere risolto con un programma simile a quello dell'esercizio precedente.
Anche qui dobbiamo leggere un input dall'utente e prendere decisioni in base al valore inserito.
La differenza è che qui le decisioni sono basate su intervalli di numeri (0-6, 7-13, ecc.), quindi useremo
istruzioni condizionali (if/else) per determinare la fase della giornata. 
Non ci sono complicazioni particolari, quindi la struttura del programma rimane semplice.
*/

package main

import "fmt"

func main() {
    var ora int
    fmt.Print("Inserisci l'ora (0-23): ")
    fmt.Scan(&ora)

    if ora < 0 || ora > 23 {
        fmt.Println("orario non valido")
    } else if ora <= 6 {
        fmt.Println("notte")
    } else if ora <= 13 {
        fmt.Println("mattino")
    } else if ora <= 18 {
        fmt.Println("pomeriggio")
    } else {
        fmt.Println("sera")
    }
}

