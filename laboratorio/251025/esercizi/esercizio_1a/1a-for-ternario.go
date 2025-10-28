package main
import "fmt"
func main() {
    /*
    In questo programma viene dichiarazione una variabile n di tipo int
    viene chiesto di inserire un numero da tastiera e successivamente viene fatto
    un ciclo for ternario stampando tanti asterischi quanto il numero che è stato inserito in input da tastiera
    e alla fine del ciclo viene messo un a capo
    */
	var n int
	fmt.Print("un numero: ")
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		fmt.Print("*")
	}
	fmt.Println()
}

