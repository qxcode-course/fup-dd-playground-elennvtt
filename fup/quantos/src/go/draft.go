package main
import "fmt"
func main() {
    var valor1, valor2, valor3 int

    fmt.Scanln(&valor1)
    fmt.Scanln(&valor2)
    fmt.Scanln(&valor3)
    if valor1 == valor2 && valor2 == valor3 {
        fmt.Println("3")
    }else if valor1 == valor2 || valor1 == valor3 || valor2 == valor3 {
        fmt.Println("2")

    }else{
        fmt.Println("0")
    }

}
