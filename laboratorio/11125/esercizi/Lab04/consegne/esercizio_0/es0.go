package main

import "fmt"

func main() {
    var carattere byte
    fmt.Print("Inserisci un carattere: ")
    fmt.Scanf("%c", &carattere) // legge un singolo byte
    fmt.Printf("Hai inserito: %c\n", carattere)

    // Stampa il byte precedente, quello stesso e quello successivo in ordine ASCII
    fmt.Printf("Byte precedente, corrente e successivo in ordine ASCII: %c %c %c\n", carattere-1, carattere, carattere+1)

    // Controllo se è una lettera tra 'A' e 'L' oppure altro
    if carattere >= 'A' && carattere <= 'L' {
        fmt.Println("A-L")
    } else {
        fmt.Println("altro")
    }

    // Legge una stringa
    var s string
    fmt.Print("Inserisci una stringa: ")
    fmt.Scan(&s)

    // Stampa la stringa in verticale, una runa per riga
    for i := 0; i < len(s); i++ {
        fmt.Printf("%c\n", s[i])
    }
}

