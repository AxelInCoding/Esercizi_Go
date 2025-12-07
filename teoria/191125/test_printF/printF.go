package main
import "fmt"
import "math"

func main() {
  var x = 153
  var f = 3.14
 
  f = math.Sqrt(2)
  fmt.Printf("<%d>\n", x)
  fmt.Printf("<%6d>\n", x)
  fmt.Printf("<%-6d>\n", x)
  fmt.Printf("<%06d>\n", x)
  fmt.Printf("<%010b>\n", x)
  fmt.Printf("<%0*b>\n", 15, x)
  fmt.Println()
  fmt.Println(f)
  fmt.Printf("<%.2f>\n", f)
  fmt.Printf("<%7.2f>\n", f)
  fmt.Printf("<%2.2f>\n", f)
  
}
