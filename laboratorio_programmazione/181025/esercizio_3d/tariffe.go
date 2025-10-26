package main
import "fmt"


func main (){
  var eta int;
  var studente bool;

  fmt.Print("Inserisci la tua età: ");
  fmt.Scan(&eta);

  fmt.Print("Sei studente (true/false): ");
  fmt.Scan(&studente);

  var tariffa float64;

  if eta < 9 {
    tariffa = 0;
  }else if eta < 18 {
    tariffa = 5;
  }else if eta < 26 {
     tariffa = 7.5; 
  }else if eta <= 65 {
      tariffa = 10;
  } else {
    tariffa = 7.5
  }

  if eta >= 18 && studente {
    tariffa = 5;
  }

  fmt.Println("Ingresso ", tariffa, " euro.");
  
}
