package main
import "fmt"


func isPrime(num int) bool {
	var divd int;

	if num < 2 {
    return false;
	}
	
	for divd = 2; divd < num; divd ++ {
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

func main () {
  var n int;
  fmt.Print("Inserisci un numero :" );
  fmt.Scan(&n);

  for i:=2; i<=n; i++ {
    if isPrime(i) && i%10 == 7 {
      fmt.Println(i);
    }
  }
}
