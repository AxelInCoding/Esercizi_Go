package main 
import "fmt"
import "bufio"
import "os"

func main() {
  scanner := bufio.NewScanner(os.Stdin)
  var lines []string;

  fmt.Println("Inserisci delle righe (CTRL+D o CTRL+Z per terminare): ")

  for scanner.Scan() {
    line := scanner.Text()
    lines = append(lines, line)
  }

  //stampa le righe in ordine inverso
  fmt.Println("\n Righe in ordine inverso: ")
  for i:=len(lines)-1; i >=0; i--{
    fmt.Println(lines[i])
  }

}



