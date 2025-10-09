package main
import "fmt"

func main (){
  const GIORNIPERANNO = 365;
  const GIORNIPERSETTIMANA = 7;

  var giorniTotali int;
  
  fmt.Print("Inserisci il numero di giorni da convertire: ");
  fmt.Scan(&giorniTotali);

  anni :=giorniTotali / GIORNIPERANNO; // gli anni interi in giorniTotali
  restoDopoAnni := giorniTotali % GIORNIPERANNO // il resto dei giorni che non costituiscono un anno intero

  settimane := restoDopoAnni / GIORNIPERSETTIMANA; //settimane intere nei giorni rimasti
  giorni := restoDopoAnni % GIORNIPERSETTIMANA; // giorni rimanenti che non formano una settimana completa
  
  fmt.Println(giorniTotali, "giorni =", anni, " anni + ", settimane, "settimane + ", giorni, "giorni");
}
