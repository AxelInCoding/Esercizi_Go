package main
import "fmt"

func isPrime(num int) bool {
	var divd int;

	if num < 2 {
    return false;
	}
	
	for divd =2; divd < num; divd ++ {
    if num % divd == 0 {
      break;
    }
	}
	
	if divd == num {
    return true;
	}else{
    return false;
	}	
}

func main() {
	var n int
	print("Inserisci un numero: ")
	fmt.Scan(&n)

	for i := 2; i <= n-2; i++ {
		if isPrime(i) && isPrime(i+2) {
			print("(", i, ", ", i+2, ")\n")
		}
	}
}



