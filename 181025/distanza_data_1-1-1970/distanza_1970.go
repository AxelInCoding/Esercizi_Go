package main
import "fmt"

//Funzione anno bisestile
func isLeapYear (year int) bool {
  return (year % 4 == 0 && year % 100 !=0) || (year % 400 == 0);
}

//Funzione che ritorna il numero di giorni in un mese di un anno specifico
func daysInMonth (month int, year int) int {
  switch month {
    case 1, 3, 5, 7, 8, 10, 12:
      return 31;
    case 4, 6, 9, 11:
      return 30;
    case 2:
      if isLeapYear(year){
        return 29;
      }else{
        return 28;
      }
    default:
      return 0;
  }
}

//Funzione che calcola i giorni dall' 1/1/1970 fino alla data specificata
func daySinceEpoch (day, month, year int) int{
  days:=0;

  //Anngiungo i giorni degli anni completi
  for y:=1970; y<year; y++ {
    if isLeapYear(y) {
      days+= 366;
    }else{
      days+= 365;
    }
  }

  //Aggiungo i giorni dei mesi completi nell'anno corrente
  for m:=1; m<month; m++ {
    days+=daysInMonth(m, year);
  }

  //Aggiungo i giorni del mese corrente
  days+= day - 1; // il primo giorno non conta

  return days;
}


func main (){
  var day, month, year int;

  fmt.Print("Inserisci il giorno: ");
  fmt.Scan(&day);
  fmt.Print("Inserisci il mese: ");
  fmt.Scan(&month);
  fmt.Print("Inserisci l'anno: ");
  fmt.Scan(&year);

  totalDays := daySinceEpoch (day, month, year);

  fmt.Println("La distanza in giorni dall' 1/1/1970 è: ", totalDays);
}
