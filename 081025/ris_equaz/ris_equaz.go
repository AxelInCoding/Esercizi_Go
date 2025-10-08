package main
import "fmt"
import "math"

func main (){

  var a, b, c float64;
  fmt.Print("Inserisci a: ");
  fmt.Scan(&a);
  fmt.Print("Inserisci b: ");
  fmt.Scan(&b);
  fmt.Print("Inserisci c: ");
  fmt.Scan(&c);

  if a == 0 {
    if b == 0 {
      if c == 0 {
        fmt.Println("Equazione indeterminata (infinite soluzioni)");
        
      }else{
      fmt.Println("Equazione impossibile (nessuna soluzione)");
      
      }
    }else{
      x:= -c / b
      fmt.Println("Equazione ineare con soluzione x= ", x);
    
    }
  }else{
    delta:= (b*b) -4* (a*c);

    if delta > 0 {
    
      x1:= (-b + math.Sqrt(delta)) / (2 * a);
      x2:= (-b - math.Sqrt(delta)) / (2 * a);
      fmt.Println("Equazione quadratica con due soluzioni reali distinte:");
      fmt.Println("x1 =", x1);
      fmt.Println("x2 =", x2);
      fmt.Println("Delta =", delta);
      
    }else if delta == 0{
    
      x:= -b / (2 * a);
      fmt.Println("Equazione quadratica con una soluzione reale doppia:");
      fmt.Println("x =", x);
      fmt.Println("Delta =", delta);
      
    }else{
      fmt.Println("Nessuna soluzione reale");
      fmt.Println("Delta =", delta);
    }
  }
  
}
