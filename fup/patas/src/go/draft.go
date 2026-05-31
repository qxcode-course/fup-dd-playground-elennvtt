package main
import "fmt"
func main() {
    var chico, cebolinha int
    fmt.Scan(&chico)
    fmt.Scan(&cebolinha)

    var n int 
    fmt.Scan(&n)

    total := 0
    var animal string

    for i := 0; i < n; i++ {
        fmt.Scan(&animal)
        switch animal {
        case "v", "c":
            total += 4 
        case "g":
            total += 2
        }
    }
    fmt.Println(total)

    dchico := abs(chico - total)
    dcebolinha := abs(cebolinha-total)

    if dchico < dcebolinha {
        fmt.Println("Chico Bento")
    }else if dcebolinha < dchico {
        fmt.Println("Cebolinha")
    }else {
        fmt.Println("empate")
    }
}
    func abs (x int) int{
        if x <0 {
            return -x
        }
        return -x
}