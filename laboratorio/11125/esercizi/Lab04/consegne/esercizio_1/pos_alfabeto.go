package main
import "fmt"

func main() {
    var c byte
    var sommaCifre int

    fmt.Println("Inserisci caratteri ASCII terminati da '.'")

    for {
        fmt.Scanf("%c", &c) // legge un carattere alla volta

        if c == '.' {
            break // termina al punto
        }

        if c >= 'a' && c <= 'z' {
            pos := int(c - 'a' + 1) // posizione nell'alfabeto
            fmt.Println(string(c) + " è la " + fmt.Sprint(pos) + "^a")
        } else if c >= '0' && c <= '9' {
            val := int(c - '0') // valore numerico
            sommaCifre += val
            fmt.Println(string(c) + " ha valore " + fmt.Sprint(val))
        } else {
            fmt.Println(string(c) + " - altro")
        }
    }

    fmt.Println("Somma delle cifre: " + fmt.Sprint(sommaCifre))
    fmt.Println("bye")
}

