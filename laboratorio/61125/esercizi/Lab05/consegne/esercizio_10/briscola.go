package main 
import "fmt"

func puntiCarta(s string) int{
  
  if len(s) != 3 {
    return -1
  }


  punti := 0
  for _, runa := range s {
    switch runa {
      case 'A':
        punti += 11
      case '3':
        punti += 10
      case 'K':
        punti += 4
      case 'Q':
        punti += 3
      case 'J':
        punti += 2
      case '7', '6', '5', '4', '2':

      default:
        return -1
    }
  }
  return punti
}

func main() {
    var mano string
    fmt.Print("Inserisci la mano (3 carte): ")
    fmt.Scanln(&mano)

    if len(mano) != 3 {
        fmt.Println("Errore: la mano deve contenere esattamente 3 caratteri.")
        return
    }

    tot := puntiCarta(mano)
    if tot == -1 {
        fmt.Println("Errore: la mano contiene simboli non validi.")
        return
    }

    fmt.Printf("Mano %s: %d punti\n", mano, tot)
}
