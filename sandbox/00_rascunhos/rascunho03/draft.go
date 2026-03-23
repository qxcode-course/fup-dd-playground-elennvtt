 package main
import "fmt"
func main() {
    var  n1, n2 int
    //entrada
    fmt.Scan(&n1)
    fmt.Scan(&n2)
    //calculo
    quociente := n1 / n2
    resto := n1 % n2 
     // saida
      fmt.Printf("%d %d/n", quociente, resto)
      

}
    
