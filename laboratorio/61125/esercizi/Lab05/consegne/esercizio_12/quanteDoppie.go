package main 
import "fmt"
import "bufio"
import "os"
import "strings"


func haDoppie (s string) bool {
    
    for i:=0; i < len(s) - 1; i++ {
      coppia := s[i : i+2]
        if strings.Count(s, coppia) > 0 && coppia[0] == coppia[1] {
          return true
        }
    }
    return false
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    count := 0

    fmt.Println("Inserisci parole (termina con Ctrl+D su nuova riga):")

    for scanner.Scan() {
        parola := strings.TrimSpace(scanner.Text())
        if parola == "" {
            continue
        }
        if haDoppie(parola) {
            count++
        }
    }

    fmt.Printf("Numero di parole con doppie: %d\n", count)
}
