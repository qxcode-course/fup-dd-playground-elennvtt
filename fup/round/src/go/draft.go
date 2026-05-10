package main
import ("fmt"
       "math"
)
func main() {
    var operacao string
    var num float64
    fmt.Scan(&operacao)
    fmt.Scan(&num)

    var resultado float64

    switch operacao {
    case "r":
        resultado = math.Round(num)
    case "f":
        resultado = math.Floor(num)
    case "c":
        resultado = math.Ceil(num)  
    }
    fmt.Println(int(resultado))
}
