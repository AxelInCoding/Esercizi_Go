package main
import(
  "fmt"
  "bufio"
  "os"
)

func main() {
  fmt.Println("scrivi parole (ctrl D per terminare)")

  scanner := bufio.NewScanner(os.Stdin)
  scanner.Split(bufio.ScanWords) // lettura parola per parola

  posizioni := make(map[string][]int)
  i := 0 //posizione corrente

  for scanner.Scan() {
    parola := scanner.Text()
    posizioni[parola] = append(posizioni[parola], i)
    i++  
  }

  if err := scanner.Err(); err != nil {
		fmt.Println("Errore lettura input:", err)
		return
	}

	fmt.Println(posizioni)
  
}  
