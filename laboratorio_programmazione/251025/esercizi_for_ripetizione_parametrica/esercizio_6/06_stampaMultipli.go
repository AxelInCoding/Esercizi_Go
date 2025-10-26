package main
import "fmt"


func main() {
	var n, p, q int

	fmt.Print("Inserisci n, p, q: ")
	fmt.Scan(&n, &p, &q)

	for i := 1; i <= n; i++ {
		if i%p == 0 && i%q != 0 {
			fmt.Print("\t", i)
		}
	}
	fmt.Println() 
}

