package main
import "fmt"

func main (){
  var r rune;
  var s string;

  fmt.Print("Inserisci un carattere: ")
  fmt.Scanf("%c", &r)

  fmt.Scanf("\n")
  
  fmt.Print("Inserisci una stringa: ")
  fmt.Scanf("%s", &s)

  pos:= -1;

  for i, runaCorrente:= range s {
    if runaCorrente == r {
      pos = i;
      break;
    }
  }
  fmt.Println("il carattere trovato nella stringa è ", string(r), "e si trova in posizione: ", pos)
}
