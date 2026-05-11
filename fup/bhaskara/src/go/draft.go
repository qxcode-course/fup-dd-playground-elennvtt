package main

import (
	"fmt"
	"math"
)
func main() {
    var a, b, c float64
    fmt.Scanln(&a)
    fmt.Scanln(&b)
    fmt.Scanln(&c)

    delta := b*b - 4*a*c

    if delta > 0 {
        x1 := (-b + math.Sqrt(delta)) / (2*a)
        x2 := (-b - math.Sqrt(delta)) / (2*a)
        fmt.Printf("%.2f\n", x1)
        fmt.Printf("%.2f\n", x2)
    }else if delta == 0{
        x := -b / (2*a)
        fmt.Printf("%.2f\n", x)
    }else {
        fmt.Println("nao ha raiz real")
    }
}
