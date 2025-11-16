package main
import "fmt"
import "io"

func main (){
  var pioggia int
  giorno := 1 //contatore giorni
  ultimoGiorno := 0 //ultimo giorno di pioggia
  
  for {
  _, err := fmt.Scan(&pioggia)
		if err == io.EOF {
			break
		} else if err != nil {
      fmt.Println("Errore nella lettura: ", err)
      return
		}

		if pioggia > 0 {
      ultimoGiorno = giorno
		}
		giorno ++
  }

  if ultimoGiorno == 0 {
    fmt.Println("Non ha piovuto iin nessun giorno.")
  } else {
    fmt.Println("Ultimo giorno di pioggia: ", ultimoGiorno)
  }

}
