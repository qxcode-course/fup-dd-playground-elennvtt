package main
import "fmt"
func main() {
    var chico, cebolinha int
    fmt.Scan(&chico)
    fmt.Scan(&cebolinha)

    var n int 
    fmt.Scan(&n)

    total := 0
    animal := make([]string, n)

    for i := 0; i < n; i++ {
        fmt.Scan(&animal[i])
        switch animal[i] {
        case "v", "c":
            total += 4 
        case "g":
            total += 2
        }
    }
    fmt.Println(total)

    dchico := abs(total - chico)
    dcebolinha := abs(total - cebolinha )

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