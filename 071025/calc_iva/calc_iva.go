package main
import "fmt"

func main () {
var scelta int;
var check int = 1;


for check==1 {
  fmt.Print(
  "Bevenuto nel convertitore, fai la tua scelta\n",
  "1. Dato IMPONIBILE in euro e ALIQUUOTA trova il PREZZO FINALE\n",
  "2. Dato il PREZZO FINALE e ALIQUOTA trova L'IMPONIBILE in euro\n",
  "3. Dato L'IMPONIBIILE e il PREZZO FINALE trova L'ALIQUOTA\n",
  "4. ESCI\n",
  );
  
  fmt.Scan(&scelta);
  
  switch scelta {
  
  case 1:
    var imponibile, aliquota float64;
    fmt.Print("Inserisci l'imponibile in euro: ");
    fmt.Scan(&imponibile);
    fmt.Print("iInserisci l'aliquota: ");
    fmt.Scan(&aliquota);
    fmt.Println();


    var percentuale float64 = (imponibile * aliquota) / 100
    fmt.Println("Il prezzo finale è: ", imponibile + percentuale);
    
  case 2:
  var prezzo_finale, aliquota float64;
    fmt.Print("Inserisci il prezzo finale in euro: ");
    fmt.Scan(&prezzo_finale);
    fmt.Print("Inserisci l'aliquota: ");
    fmt.Scan(&aliquota);
    fmt.Println();

    fmt.Println("I'imponibile è: ", (100 * prezzo_finale) / (100 + aliquota));

  case 3:
    var prezzo_finale, imponibile float64;
    fmt.Print("Inserisci il prezzo finale in euro: ");
    fmt.Scan(&prezzo_finale);
    fmt.Print("Inserisci l'imponibile in euro: ");
    fmt.Scan(&imponibile);
    fmt.Println();
    
    fmt.Println("L'aliquota è: ", ((100 * prezzo_finale) / imponibile) - 100);

  
  case 4:
    fmt.Println("Grazie per aver utilizzato il programma!");
    check = 0;
    
  default:
    fmt.Println("Comando errato!");
  }
};

}//fine MAIN

