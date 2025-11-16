package main
import "fmt"
import "unicode"
import "bufio"
import "os"

/*
In questo esercizio prendiamo in considerazione solo caratteri ASCII (byte).
In un file funzioni.go definire le seguenti funzioni:

nomefile: funzioni.go

/*
- hasUpper(s string) bool
La funzione riceve una stringa come parametro e restituisce 
true se la stringa contiene almeno una lettera latina maiuscola (tra 'A' e 'Z'), 
false altrimenti.
*/

func hasUpper (s string) bool {
  for _, runa := range s {
    if unicode.IsUpper(runa){
      return true
    }
  }
  return false
}

/*
- firstUpper(s string) int
La funzione riceve una stringa come parametro e restituisce
 la posizione della prima lettera latina maiuscola (tra 'A' e 'Z'),
  se la stringa ne contiene almeno una, -1 altrimenti.
*/

func firstUpper (s string) int {
  for i, runa := range s {
    if unicode.IsUpper(runa){
      return i
    }
  }
  return -1
}

/*
- lastUpper(s string) int
La funzione riceve una stringa come parametro e restituisce la
 posizione dell'ultima lettera latina 
maiuscola (tra 'A' e 'Z'), se la stringa
 ne contiene almeno una, -1 altrimenti.
*/

func lastUpper (s string) int {

  last := -1
    
  for i, runa := range s {
    if unicode.IsUpper(runa) {
      last = i
    }
  }
  return last
}

/*
- countDigitsLettersOthers(s string) int int int
La funzione riceve una stringa come parametro, 
conta quante cifre, quante lettere e quanti altri caratteri 
(né cifre né lettere) contiene e restituisce i tre risultati 
in questo ordine.
*/

func countDigitsLettersOthers (s string) (int, int, int) {

  digits := 0 
  letters := 0
  others := 0
  
  for _, runa := range s {
  
    if unicode.IsDigit(runa) {
      digits++
    } else if unicode.IsLetter(runa) {
      letters++
    } else {
      others++
    }
  }
  return digits, letters, others
}

/*
- isPalindrome(s string) bool
La funzione riceve una stringa come parametro 
e restituisce true se la stringa è palindroma,
 e false altrimenti. Una parola è palindroma se
  può essere letta  sia da sinistra a destra che da destra a sinistra. Ad esempio "osso" e "ingegni" sono palindrome, ma anche "" (la stringa vuota) e "t" (le stringhe di un solo carattere).
*/

func isPalindrome (s string) bool {
  
  n := len(s)

  if n<=1 {
    return true
  }

  for i:=0; i < n/2; i++ {
    if s[i] == s[n-1-i]{
      return true;
    }
  }
  return false
}


/*
Scrivete infine un main che legge una stringa da standard input, usa le funzioni
qui sopra per determinare se la stringa letta contiene maiuscole, in che posizione
è la prima maiuscola, in che posizione è l'ultima maiuscola, quante cifre, lettere e
altri caratteri contiene, se è palindroma, e stampa i risultati 
("ha maiuscole" / "non ha maiuscole", "prima maiuscola in posizione ...", “palindroma" / "non palindroma", ecc.).*/
func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Inserisci una stringa: ")
	
	if scanner.Scan() { 
		input := scanner.Text()
	
		if hasUpper(input) {
      fmt.Println("(1) La stringa contiene almeno una maiuscola")
		} else {
      fmt.Println("(1) La stringa non ha maiuscole.")
		}
		
		first := firstUpper(input)
		last := lastUpper(input)

		if first != 1 {
      fmt.Println("(2) Prima maiuscola in posizione: ", first)
		} else {
      fmt.Println("(2) Non ci sono maiuscole.")
		}

		if last != 1 {
      fmt.Println("(2) Ultima maiuscola in posizione: ", last)
		}


		digits, letters, others := countDigitsLettersOthers(input)
		fmt.Println("(3) Cifre: ", digits)
		fmt.Println("(3) lettere: ", letters)
		fmt.Println("(3) Altri caratteri: ", others)


		if isPalindrome(input) {
      fmt.Println("(4) La stirnga è Palindroma.")
		} else {
      fmt.Println("(4) la stringa NON è Palindroma.")
		}
	}

}
