package main
import "fmt"
import "strconv"

func giorniInMese(mese int) int {
    switch mese {
    case 1, 3, 5, 7, 8, 10, 12:
        return 31
    case 4, 6, 9, 11:
        return 30
    case 2:
        return 28
    default:
        return 0 // Non dovrebbe mai capitare dato che assumiamo 1-12
    }
}

func main() {
    var data string
    fmt.Print("Inserisci la data (gg-mm-aaaa): ")
    fmt.Scanln(&data)

    // Estrarre il mese come sottostringa
    meseStr := data[3:5] 
    mese, err := strconv.Atoi(meseStr)
    
    if err != nil {
        fmt.Println("Formato data non valido")
        return
    }

    giorni := giorniInMese(mese)
    fmt.Printf("Il mese %d ha %d giorni\n", mese, giorni)
}
