package main
import "fmt"
func main() {
    var p1, p2, p3 float64
    var v1, v2, v3 float64
    var din float64
   
   //entrada
   //quantidade
   fmt.Scan(&p1)
   fmt.Scan(&p2)
   fmt.Scan(&p3)
   // valores
   fmt.Scan(&v1)
   fmt.Scan(&v2)
   fmt.Scan(&v3)
   fmt.Scan(&din)

   //calculo
   v1 = v1 * p1
   v2 = v2 * p2
   v3 = v3 * p3
   //saida 
 var troco float64 = din - (v1 + v2 + v3)
   fmt.Printf("%.2f\n", troco)
}




