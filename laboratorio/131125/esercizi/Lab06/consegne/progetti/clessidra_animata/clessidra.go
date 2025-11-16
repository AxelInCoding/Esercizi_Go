package main 
import (
  "bufio"
  "fmt"
  "os"
  "os/exec"
  "strings"
  "strconv"
  "time"
)

//rigaClessidra costruisce una riga intera della clessidra
func rigaClessidra(length int, sand bool, sandChar byte) string {
  if sand {
    return strings.Repeat(string(sandChar), length)
  } 
  return strings.Repeat(" ", length)
}

//clessidra stampa la clessidra a un dato tempo
func clessidra(height int, t int, char byte) {
  base := height * 2 + 1

  //parte superiore
  for i := 0; i < height; i++ {
    //calcolo larghezza interna
    inner := base - 2 * (i + 1)
    sandLayers := height - t //sabbia rimasta sopra
    sand := sandLayers > i

    fmt.Print(strings.Repeat(" ", i))
    fmt.Print("\\")
    fmt.Print(rigaClessidra(inner, sand, char))
    fmt.Println("/")
  }

  //parte inferiore
  for i := height - 1; i >= 0; i-- {
    inner := base -2 * (i + 1)
    fallen := t //sabbia scesa
    sand := fallen > (height - 1 - i)

    fmt.Print(strings.Repeat(" ", i))
    fmt.Print("/")
    fmt.Print(rigaClessidra(inner, sand, char))
    fmt.Println("\\")
  }
}

//funzione che fa il clear dello schermo
func cancella() {
  cmd := exec.Command("clear")
  cmd.Stdout = os.Stdout
  cmd.Run()
}

//funzione che attende
func attendi(n_seconds float64) {
  time.Sleep(time.Duration(n_seconds) * time.Second)
}

func main () {
  scanner := bufio.NewScanner(os.Stdin)
  
  fmt.Print("Inserisci il numero di secondi: ")
  scanner.Scan()
  secStr := scanner.Text()

  sec, _ := strconv.Atoi(secStr)

  for t := 0; t <= sec; t++ {
    cancella()
    clessidra(sec, t, '*')
    attendi(1)
  }
}
