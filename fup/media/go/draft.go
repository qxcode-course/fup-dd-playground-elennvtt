package main
import "fmt"
func main() {
    var num1, num2 float64
    fmt.Scan(&num1)
    fmt.Scan(&num2)
    
    media := (num1 + num2) / 2

    fmt.Printf("%.1f\n", media)


}
