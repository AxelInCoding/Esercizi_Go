/*
una stringa in formato CSV (del tipo: <nome>,<cognome>,<matricola>), 
con un numero non noto a priori di valori,

ne estrae i valori salvandoli in una slice e stampa la slice;
dalla stessa stringa estrae i valori compresa la virgola 
separatore salvandoli in una slice e stampa la slice;


una frase tra "" di parole separate anche da più spazi 
e tab, ne estrae le parole salvandole in una slice e 
stampa la slice;
un numero intero non negativo e stampa la slice dei 
codici ASCII delle cifre che lo formano, nello stesso 
ordine da sinistra a destra;
una stringa di lettere accentate e stampa la slice 
dei codici unicode delle lettere che la formano;
un orario nel formato h:m:s, ne estrae ore,
 minuti e secondi salvandoli in una slice e stampa la slice.
*/

package main 
import (
	"fmt"
	"os"
	"strings"
)

func main() {
  if len(os.Args) != 6 {
    fmt.Println("Uso: ./strings2slice <csv> <frase> <numero> <accentate> <orario>")
		return
  } 

  // 1)
  csv := os.Args[1]

  //valori senza virgola
  valori := strings.Split(csv, ",")
  fmt.Println(valori)

  //valori con virgola
  valoriConSep := strings.SplitAfter(csv, ",")
  //Rimuovo l'eventuale stringa vuota finale
  if valoriConSep[len(valoriConSep)-1] == "" {
    valoriConSep = valoriConSep[:len(valoriConSep)-1]
  }
  fmt.Println(valoriConSep)

  // 2)
  frase := os.Args[2]
  parole := strings.Fields(frase)
  fmt.Println(parole)

  // 3)
  numero := os.Args[3]
  byteCodes := []byte(numero)
  var codiciASCII []int
  for _, b := range byteCodes {
		codiciASCII = append(codiciASCII, int(b))
	}
	fmt.Println(codiciASCII)

	// 4)
	accentate := os.Args[4]
	runes := []rune(accentate)
	var codiciUnicode []int
	for _, r := range runes {
		codiciUnicode = append(codiciUnicode, int(r))
	}
	fmt.Println(codiciUnicode)

	//5)
	orario := os.Args[5]
	parts := strings.Split(orario, ":")
	fmt.Println(parts)
  
}
