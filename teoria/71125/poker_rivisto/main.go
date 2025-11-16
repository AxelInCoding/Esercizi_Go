package main 
import "fmt"

func main () {
  var m mazzo
  m = mazzoIntero(5)
  mescola(m)

  for_, c := range m {
    fmt.Println(String(c))
  }

  var num int
  fmt.Print("Quante vcarte vuoi cambiare? [0-3]")
  for {
    fmt.Scan(&num)
    if num <0 || num>3 {
    fmt.Print("*** Valore non consentito")
    } else {
      break
    }
  }


  var indiciDaCambiare [] int
  indiciDaCambiare = make ([]int, num) //L'utente vuole cambiare num
  for i:=0; i < num; i++{
    fmt.Print("Indice carta da cambiar : ")
    fmt.Scan(&indiciDaCambiare[i])
  }

  carteNuove := servi(m[5:], num)
  for i:=0; i < num; i++ {
    mano[indiciDaCambiare[i]]= 
  }
}
